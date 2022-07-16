package context

import (
	"context"
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/repository"
	"github.com/carlosrodriguesf/dfile/src/tool/dbfile"
)

type (
	Context interface {
		context.Context

		DBFile() dbfile.DBFile
		SetDBFile(dbFile dbfile.DBFile)

		Load(db *sql.DB)

		App() app.Container
	}

	contextImpl struct {
		context.Context

		dbFile       dbfile.DBFile
		apps         app.Container
		repositories repository.Container
	}
)

func New() Context {
	return &contextImpl{
		Context: context.Background(),
	}
}

func (c *contextImpl) DBFile() dbfile.DBFile {
	return c.dbFile
}

func (c *contextImpl) App() app.Container {
	return c.apps
}

func (c *contextImpl) Load(db *sql.DB) {
	c.apps = app.NewContainer(repository.NewContainer(db))
}

func (c *contextImpl) SetDBFile(dbFile dbfile.DBFile) {
	c.dbFile = dbFile
}
