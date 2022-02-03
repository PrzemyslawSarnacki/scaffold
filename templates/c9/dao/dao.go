package dao

import (
	uuid "github.com/satori/go.uuid"
	base "gitlab.com/clowd9/platform/base/db"
)


type {{.ProjectName}}Dao interface {
	Get{{.ProjectName}}(id uuid.UUID) (*base.{{.ProjectName}}, error)
}
