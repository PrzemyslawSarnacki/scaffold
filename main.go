package main

import (
	"os"
	"path/filepath"

	"github.com/PrzemyslawSarnacki/scaffold/scaffold"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	var serviceType string
	var projectName string
	var templateName string
	var destinationDir string

	app := cli.NewApp()
	app.Version = "1.0.0-rc"
	app.Usage = "Generate scaffold project layout for Go."
	app.Commands = []*cli.Command{
		{
			Name:        "init",
			Aliases:     []string{"i"},
			Description: "Generate scaffold project layout",
			Usage:       "ex. scaffold init --destination ../ --template c9 --servicetype module --name webscraper",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "destination",
					Aliases: []string{"d", "dest"},
					Usage:       "Specify destination of generated project",
					Value:       "",
					Destination: &destinationDir,
				},
				&cli.StringFlag{
					Name:        "template",
					Aliases: []string{"t"},
					Usage:       "Specify template type ex. c9, webapp",
					Value:       "c9",
					Destination: &templateName,
				},
				&cli.StringFlag{
					Name:        "servicetype",
					Aliases: []string{"service","type", "s"},
					Usage:       "Specify service type",
					Value:       "module",
					Destination: &serviceType,
				},
				&cli.StringFlag{
					Name:        "name",
					Aliases: []string{"n"},
					Usage:       "Specify name for the project",
					Value:       "project",
					Destination: &projectName,
				},
			},
			Action: func(c *cli.Context) error {
				// currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
				currDir, err := getDirectory(destinationDir)
				if err != nil {
					log.Errorf("error while getting Directory: %v", err)
					return err
				}
				currDir = filepath.Join(currDir, projectName)
				err = os.Mkdir(currDir, os.ModePerm)
				if err != nil {
					log.Error("error while creating dir: %v",err)
					return err
				}
				err = scaffold.New(false).Generate(currDir, projectName, templateName, serviceType)
				if err == nil {
					log.Printf("Successfully created project: %s from template: %s", projectName, templateName)
				}
				return err
			},
		},
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Description: "List all available templates",
			Usage:       "ex. scaffold list",
			Action: func(c *cli.Context) error {
				templateStruct := []struct {
					name string
					desc	string
					
				}{
					{
						name: "template",
						desc: "standard scaffold template",
					},
					{
						name: "c9",
						desc: "c9 starter microservice template",
					},
				}
				log.Info("Available templates:\n")
				for i, t := range templateStruct {
					log.Infof("%d. %s - %s", i+1, t.name, t.desc)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// if directory not specified get current directory
func getDirectory(path string) (string, error) {
	if path != "" {
		currentDir, err := filepath.Abs(filepath.Dir(path))
		if err != nil {
			log.Errorf("error while getting current path: %v", err)
			return "", err
		}
		return currentDir, nil
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Errorf("error while getting current working directory: %v", err)
		return "", err
	}
	return currentDir, nil
}
