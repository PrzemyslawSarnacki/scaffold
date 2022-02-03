package svc

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"gitlab.com/clowd9/platform/micro/service/dependency"
	"gitlab.com/clowd9/platform/micro/service/provider"
)

func Cli(pr dependency.DependencyProvider) ([]*cli.Command, error) {

	c9int, err := pr.GetDependency(dependency.DependencyC9)
	if err != nil {
		return nil, err
	}
	c9 := c9int.(provider.CliClient)

	return append([]*cli.Command{}, []*cli.Command{
		{
			Name:  "{{.ProjectName}}",
			Usage: "{{.ProjectName}}",
			Subcommands: []*cli.Command{
				{
					Name:        "hc",
					Usage:       "{{.ProjectName}} hc",
					Description: "health check",
					Action: func(c *cli.Context) error {
						return c9.Call(ServiceName, "HealthCheck", `{}`)
					},
				},
			},
		},
	}...), nil
}
