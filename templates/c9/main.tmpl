package main

import (
	basedb "gitlab.com/clowd9/platform/base/db/cockroach"
	"gitlab.com/clowd9/platform/micro/service"
	"gitlab.com/clowd9/platform/micro/service/dependency"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Create new service
	conn, err := basedb.NewConnection()
	if err != nil {
		log.Fatal(err)
		return
	}

	dpr := dependency.NewProvider()
	spr := CreateServiceProvider(dpr)
	app := service.NewService(spr.GetServiceName())
	dpr.AddDependency(dependency.DependencyDatabase, conn.GetDb())
	err = spr.InitService()(app)
	if err != nil {
		log.Fatal(err)
	}
	// Run service
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
