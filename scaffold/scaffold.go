package scaffold

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	GoScaffoldPath = "src/scaffold"
)

var GoPath string

func init() {
	GoPath = os.Getenv("GOPATH")
	if GoPath == "" {
		panic("cannot find $GOPATH environment variable")
	}
}

type scaffold struct {
	debug bool
}

type data struct {
	AbsGenProjectPath string // The Abs Gen Project Path
	ProjectPath       string //The Go import project path (eg:github.com/fooOrg/foo)
	ServiceType       string //The project name which want to generated
	ServiceTypeTitle  string //The project name which want to generated
	ProjectName       string //The project name which want to generated
	ProjectNameTitle  string //Capitalized first letter
	Quit              string
}

func New(debug bool) *scaffold {
	return &scaffold{debug: debug}
}

func (s *scaffold) Generate(path, projectName, templateName, serviceType string) error {
	genAbsDir, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	// projectName := filepath.Base(genAbsDir)
	//TODO: have to check path MUST be under the $GOPATH/src folder
	goProjectPath := strings.TrimPrefix(genAbsDir, filepath.Join(GoPath, "src")+string(os.PathSeparator))

	d := data{
		AbsGenProjectPath: genAbsDir,
		ProjectPath:       goProjectPath,
		ProjectName:       projectName,
		ProjectNameTitle:  strings.Title(projectName),
		ServiceType:       serviceType,
		ServiceTypeTitle:  strings.Title(serviceType),
		Quit:              "<-quit",
	}

	if err := s.genFromTemplate(getTemplateSets(templateName), d); err != nil {
		return err
	}

	if err := s.genStaticFiles(getStaticTemplateSets(templateName), d); err != nil {
		return err
	}
	return nil
}

func (s *scaffold) genFromTemplate(templateSets []templateSet, d data) error {
	for _, tmpl := range templateSets {
		if err := s.tmplExec(tmpl, d); err != nil {
			return err
		}
	}
	return nil
}
func (s *scaffold) genStaticFiles(templateSets []templateSet, d data) error {
	for _, tmpl := range templateSets {
		if err := s.genStatic(tmpl, d); err != nil {
			return err
		}
	}
	return nil
}

func unescaped(x string) interface{} { return template.HTML(x) }

func (s *scaffold) tmplExec(tmplSet templateSet, d data) error {
	tmpl := template.New(tmplSet.templateFileName)
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": unescaped})
	tmpl, err := tmpl.ParseFiles(tmplSet.templateFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}

	relateDir := filepath.Dir(tmplSet.genFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(tmplSet.genFilePath))
	distAbsFilePath := filepath.Join(d.AbsGenProjectPath, distRelFilePath)

	s.debugPrintf("distRelFilePath:%s\n", distRelFilePath)
	s.debugPrintf("distAbsFilePath:%s\n", distAbsFilePath)

	if err := os.MkdirAll(filepath.Dir(distAbsFilePath), os.ModePerm); err != nil {
		return pkgErr.WithStack(err)
	}

	dist, err := os.Create(distAbsFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	defer dist.Close()

	log.Printf("Create %s\n", distRelFilePath)
	return tmpl.Execute(dist, d)
}

func (s *scaffold) genStatic(tmplSet templateSet, d data) error {
	path := filepath.Join(tmplSet.templateFilePath)
	src, err := os.Open(path)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	defer src.Close()

	// basepath := filepath.Join(GoPath, GoScaffoldPath, "templates", "c9")
	relateDir := filepath.Dir(tmplSet.genFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(tmplSet.genFilePath))
	distAbsFilePath := filepath.Join(d.AbsGenProjectPath, distRelFilePath)

	if err := os.MkdirAll(filepath.Dir(distAbsFilePath), os.ModePerm); err != nil {
		return pkgErr.WithStack(err)
	}

	dist, err := os.Create(distAbsFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	defer dist.Close()

	if _, err := io.Copy(dist, src); err != nil {
		return pkgErr.WithStack(err)
	}

	log.Printf("Create %s \n", distRelFilePath)
	return nil
}

func (s *scaffold) debugPrintf(format string, a ...interface{}) {
	if s.debug {
		log.Printf(format, a...)
	}
}
