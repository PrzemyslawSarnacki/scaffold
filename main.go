package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "github.com/catchplay/scaffold/scaffold"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0-rc"
	app.Usage = "Generate scaffold project layout for Go."
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Description: "Generate scaffold project layout",
			Usage:   "ex. scaffold init destination path",
			Action: func(c *cli.Context) error {
				// currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
				currDir, err := getDirectory(os.Args)
				if err != nil {
					return err
				}
				fmt.Println(currDir)
				fmt.Println(os.Args)
				// err = scaffold.New(false).Generate(currDir)
				// //fmt.Printf("error:%+v\n", err)
				// if err == nil {
				// 	fmt.Println("Success Created. Please excute `make up` to start service.")
				// }

				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
// if directory not specified get current directory
func getDirectory(args []string) (string, error) {
	if len(args) == 2 {
		args = append(args, "")
	}
	
	if args[2] != ""{
		currentDir, err := filepath.Abs(filepath.Dir(args[2])) 
		if err != nil {
			return "", err
		}
		return currentDir, nil
	}
	
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return currentDir, nil  
}