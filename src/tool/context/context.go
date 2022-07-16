package context

import (
	"context"
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/repository"
)

type (
	Context interface {
		context.Context

		Load(db *sql.DB)
		App() app.Container
	}

	contextImpl struct {
		context.Context

		apps         app.Container
		repositories repository.Container
	}
)

func New() Context {
	return &contextImpl{
		Context: context.Background(),
	}
}

func (c *contextImpl) App() app.Container {
	return c.apps
}

func (c *contextImpl) Load(db *sql.DB) {
	c.apps = app.NewContainer(repository.NewContainer(db))
}
