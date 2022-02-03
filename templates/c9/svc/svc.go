package svc

import (
	"github.com/urfave/cli/v2"
	"gitlab.com/clowd9/dev/sirius/core/{{.ProjectName}}/handler"
	"gitlab.com/clowd9/dev/sirius/core/{{.ProjectName}}/hmodule"
	"gitlab.com/clowd9/dev/sirius/core/{{.ProjectName}}/hmodule/config"
	pb "gitlab.com/clowd9/dev/sirius/core/{{.ProjectName}}/pb"
	"gitlab.com/clowd9/platform/micro/client"
	"gitlab.com/clowd9/platform/micro/service"
	"gitlab.com/clowd9/platform/micro/service/dependency"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const ServiceName = "core-{{.ProjectName}}"

type ServiceInitType func(provider dependency.DependencyProvider) func(app *service.Service) error

type ServiceProvider struct {
	DependencyProvider dependency.DependencyProvider
}

func NewServiceProvider(provider dependency.DependencyProvider) *ServiceProvider {
	return &ServiceProvider{
		DependencyProvider: provider,
	}
}

func (sp *ServiceProvider) InitService() func(app *service.Service) error {
	return func(app *service.Service) error {

		h := &handler.{{.ProjectName}}ModuleService{
		}

		app.AddHandler(h)
		// Register the healtcheck server
		healthpb.RegisterHealthServer(app.Server(), h)
		// Register the validator server
		pb.Register{{.ProjectName}}CoreServiceServer(app.Server(), h)

		return nil
	}
}

func (sp *ServiceProvider) GetServiceClient(cl *client.Client) interface{} {
	return pb.New{{.ProjectName}}CoreServiceClient(cl.Dial())
}

func (sp *ServiceProvider) GetServiceProtoName() string {
	return "module{{.ProjectName}}.{{.ProjectName}}ModuleService"
}

func (sp *ServiceProvider) GetServiceName() string {
	return ServiceName
}

func (sp *ServiceProvider) GetCli() ([]*cli.Command, error) {
	return Cli(sp.DependencyProvider)
}
