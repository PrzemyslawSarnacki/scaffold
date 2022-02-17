package scaffold

import (
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

type templateEngine struct {
	Templates       []templateSet
	StaticTemplates []templateSet
	currDir         string
	basePath        string
}

func (templEngine *templateEngine) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if ext := filepath.Ext(path); ext == ".tmpl" {
		templateFileName := filepath.Base(path)

		genFileBaeName := strings.TrimSuffix(templateFileName, ".tmpl") + ".go"
		err := templEngine.appendTemplate(path, genFileBaeName, false)
		if err != nil {
			return err
		}
	} else if isStatic(path) {
		templateFileName := filepath.Base(path)

		err := templEngine.appendTemplate(path, templateFileName, true)
		if err != nil {
			return err
		}
	} else if mode := f.Mode(); mode.IsRegular() {
		templateFileName := filepath.Base(path)

		err := templEngine.appendTemplate(path, templateFileName, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func (templEngine *templateEngine) appendTemplate(path string, templateFileName string, static bool) error {
	targpath := filepath.Join(filepath.Dir(path), templateFileName)
	genFileBasePath, err := filepath.Rel(templEngine.basePath, targpath)
	if err != nil {
		return pkgErr.WithStack(err)
	}

	templ := templateSet{
		templateFilePath: path,
		templateFileName: templateFileName,
		genFilePath:      filepath.Join(templEngine.currDir, genFileBasePath),
	}
	if static {
		templEngine.StaticTemplates = append(templEngine.StaticTemplates, templ)
	} else {
		templEngine.Templates = append(templEngine.Templates, templ)
	}
	return nil
}

func isStatic(path string) bool {
	staticExtenstions := []string{".png", ".gif"}
	ext := filepath.Ext(path)
	for _, staticExtenstion := range staticExtenstions {
		if ext == staticExtenstion {
			return true
		}
	}
	return false
}
