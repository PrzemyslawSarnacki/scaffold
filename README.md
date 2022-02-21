# scaffold

[![Build Status](https://travis-ci.org/PrzemyslawSarnacki/scaffold.svg)](https://travis-ci.org/PrzemyslawSarnacki/scaffold)
[![codecov](https://codecov.io/gh/PrzemyslawSarnacki/scaffold/branch/master/graph/badge.svg)](https://codecov.io/gh/PrzemyslawSarnacki/scaffold)
[![Go Report Card](https://goreportcard.com/badge/github.com/PrzemyslawSarnacki/scaffold)](https://goreportcard.com/report/github.com/PrzemyslawSarnacki/scaffold)
[![GoDoc](https://godoc.org/github.com/PrzemyslawSarnacki/scaffold?status.svg)](https://godoc.org/github.com/PrzemyslawSarnacki/scaffold)

This project is a fork of github.com/catchplay/scaffold . I reorganized structure and added more functionalities to that project and made it more customizable for my own usage than the original one. Now it can generate more than only one template type. 
Scaffold generates starter Go project based on a certain template. It let's you focus on implementation of logic instead of creating it by hand or copying stuff. 

[![asciicast](https://asciinema.org/a/MA0ppdKfZSEl64cskUnqfsSiH.svg)](https://asciinema.org/a/MA0ppdKfZSEl64cskUnqfsSiH?autoplay=1&speed=2)

## Available Code Templates

Template Name  | Template Description
------------- | -------------
template  | Default scaffold template from original `github.com/catchplay/scaffold` project
c9  | Starter for sample c9 microservice
modern  | Go application boilerplate and example applying modern practices `github.com/sagikazarmark/modern-go-application`
fiber  | Fiber web framework starter boilerplate from `github.com/gofiber/boilerplate`

The following is `template` Go project layout scaffold generated:

```
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   ├── config.yml
│   ├── database.go
│   ├── http.go
│   └── release.go
├── docker-compose.yml
├── model
│   └── model.go
└── web
    ├── routes.go
    ├── server.go
    └── version.go
```

## Installation

 Download scaffold by using:
```sh
$ go get -u github.com/PrzemyslawSarnacki/scaffold
```

## Create a new project

1. Going to your new project folder:
```sh
# change to project directory
$ cd $GOPATH/src/path/to/project
```

2. Run `scaffold init` in the new project folder (in this example for fiber template):

```sh
$ scaffold init --name projectname --template fiber
```

3. That will generate a whole new starter project files like:

```
Create .dockerignore                         
Create Dockerfile                            
Create LICENSE                               
Create Makefile                              
Create README.md                             
Create app.go                                
Create database/database.go                  
Create go.mod                                
Create go.sum                                
Create handlers/handlers.go                  
Create models/user.go                        
Create static/private/404.html               
Create static/public/css/style.css           
Create static/public/img/logo.svg            
Create static/public/index.html              
Create static/public/js/app.js               
Create static/public/img/meme.png            
Successfully created project: project from template: fiber 

```

4. And you can run the new project by using:
```sh
$ make run-local 
```

5. You are ready to code your billion dollar startup :D