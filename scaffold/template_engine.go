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

func (t *templateEngine) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if ext := filepath.Ext(path); ext == ".tmpl" {
		templateFileName := filepath.Base(path)

		genFileName := strings.TrimSuffix(templateFileName, ".tmpl") + ".go"
		targpath := filepath.Join(filepath.Dir(path), genFileName)
		err := t.appendTemplate(path, false, templateFileName, targpath)
		if err != nil {
			return err
		}
	} else if isStatic(path) {
		templateFileName := filepath.Base(path)

		err := t.appendTemplate(path, true, templateFileName)
		if err != nil {
			return err
		}
	} else if mode := f.Mode(); mode.IsRegular() {
		templateFileName := filepath.Base(path)

		err := t.appendTemplate(path, false, templateFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *templateEngine) appendTemplate(path string, static bool, paths ...string) error {
	if len(paths) == 0 {
		return pkgErr.New("not enough arguments of paths in appendTemplate function")
	}
	var targpath string
	templateFileName := paths[0]
	targpath = filepath.Join(filepath.Dir(path), templateFileName)
	if len(paths) == 2 {
		targpath = paths[1]
	}

	genFileBasePath, err := filepath.Rel(t.basePath, targpath)
	if err != nil {
		return pkgErr.WithStack(err)
	}

	templ := templateSet{
		templateFilePath: path,
		templateFileName: templateFileName,
		genFilePath:      filepath.Join(t.currDir, genFileBasePath),
	}
	if static {
		t.StaticTemplates = append(t.StaticTemplates, templ)
	} else {
		t.Templates = append(t.Templates, templ)
	}
	return nil
}

func isStatic(path string) bool {
	staticExtenstions := []string{".png", ".gif", ".yaml", ".j2", ".yml", ".conf"}
	ext := filepath.Ext(path)
	for _, staticExtenstion := range staticExtenstions {
		if ext == staticExtenstion {
			return true
		}
	}
	return false
}
