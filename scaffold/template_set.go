package scaffold

import (
	"path/filepath"
)

type templateSet struct {
	templateFilePath string
	templateFileName string
	genFilePath      string
}

func getTemplateSets(templateName string) []templateSet {
	tt := templateEngine{
		basePath:  filepath.Join(GoPath, GoScaffoldPath, "templates", templateName),
	}
	templatesFolder := filepath.Join(GoPath, GoScaffoldPath, "templates", templateName)
	//fmt.Printf("walk:%s\n", templatesFolder)
	filepath.Walk(templatesFolder, tt.visit)
	return tt.Templates
}