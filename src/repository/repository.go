package repository

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/repository/file"
	"github.com/carlosrodriguesf/dfile/src/repository/path"
)

type (
	Container interface {
		Path() path.Repository
		File() file.Repository
	}
	container struct {
		db   *sql.DB
		path path.Repository
		file file.Repository
	}
)

func NewContainer(db *sql.DB) Container {
	return &container{
		db: db,
	}
}

func (c container) Path() path.Repository {
	if c.path == nil {
		c.path = path.NewRepository(c.db)
	}
	return c.path
}

func (c container) File() file.Repository {
	if c.file == nil {
		c.file = file.NewRepository(c.db)
	}
	return c.file
}
