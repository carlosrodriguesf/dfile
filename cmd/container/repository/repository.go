package repository

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/pkg/repository/file"
	"github.com/carlosrodriguesf/dfile/pkg/repository/path"
)

type (
	Params struct {
		DB *sql.DB
	}
	Container struct {
		Path path.Repository
		File file.Repository
	}
)

func NewContainer(params Params) *Container {
	return &Container{
		Path: path.NewRepository(params.DB),
		File: file.NewRepository(params.DB),
	}
}
