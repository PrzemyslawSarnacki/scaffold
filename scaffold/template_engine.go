package scaffold

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

type templateEngine struct {
	Templates []templateSet
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
		src, err := os.Open(path)
		if err != nil {
			return pkgErr.WithStack(err)
		}
		defer src.Close()

		basepath := templEngine.basePath
		distRelFilePath, err := filepath.Rel(basepath, filepath.Join(filepath.Dir(path), filepath.Base(path)))
		if err != nil {
			return pkgErr.WithStack(err)
		}

		distAbsFilePath := filepath.Join(templEngine.currDir, distRelFilePath)

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

		// log.Printf("Create %s \n", distRelFilePath)
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