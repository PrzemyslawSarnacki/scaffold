package scaffold

import (
	// "io"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

type templateEngine struct {
	Templates []templateSet
	StaticTemplates []templateSet
	currDir   string
	basePath   string
}

func (templEngine *templateEngine) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if ext := filepath.Ext(path); ext == ".tmpl" {
		templateFileName := filepath.Base(path)

		genFileBaeName := strings.TrimSuffix(templateFileName, ".tmpl") + ".go"
		genFileBasePath, err := filepath.Rel(templEngine.basePath, filepath.Join(filepath.Dir(path), genFileBaeName))
		if err != nil {
			return pkgErr.WithStack(err)
		}

		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      filepath.Join(templEngine.currDir, genFileBasePath),
		}

		templEngine.Templates = append(templEngine.Templates, templ)

	} else if ext := filepath.Ext(path); ext == ".png" || ext == ".gif" {
		templateFileName := filepath.Base(path)

		basepath := templEngine.basePath
		targpath := filepath.Join(filepath.Dir(path), templateFileName)
		genFileBasePath, err := filepath.Rel(basepath, targpath)
		if err != nil {
			return pkgErr.WithStack(err)
		}

		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      filepath.Join(templEngine.currDir, genFileBasePath),
		}

		templEngine.StaticTemplates = append(templEngine.StaticTemplates, templ)
	} else if mode := f.Mode(); mode.IsRegular() {
		templateFileName := filepath.Base(path)

		basepath := templEngine.basePath
		targpath := filepath.Join(filepath.Dir(path), templateFileName)
		genFileBasePath, err := filepath.Rel(basepath, targpath)
		if err != nil {
			return pkgErr.WithStack(err)
		}

		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      filepath.Join(templEngine.currDir, genFileBasePath),
		}

		templEngine.Templates = append(templEngine.Templates, templ)
	}

	return nil
}