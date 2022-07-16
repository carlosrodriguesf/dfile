package repository

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/repository/path"
)

type (
	Container interface {
		Path() path.Repository
	}
	container struct {
		db   *sql.DB
		path path.Repository
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
	return nil
}
