package scaffold

import (
	"path/filepath"
)

type templateSet struct {
	templateFilePath string
	templateFileName string
	genFilePath      string
}

func getTemplateSets() []templateSet {
	tt := templateEngine{}
	templatesFolder := filepath.Join(GoPath, GoScaffoldPath, "templates/c9/")
	//fmt.Printf("walk:%s\n", templatesFolder)
	filepath.Walk(templatesFolder, tt.visit)
	return tt.Templates
}