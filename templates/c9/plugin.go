package main

import (
	"gitlab.com/clowd9/dev/sirius/module/{{.ProjectName}}/svc"
	"gitlab.com/clowd9/platform/micro/service/dependency"
	"gitlab.com/clowd9/platform/micro/service/provider"
)

//nolint
func CreateServiceProvider(dpr dependency.DependencyProvider) provider.Service {
	return svc.NewServiceProvider(dpr)
}
