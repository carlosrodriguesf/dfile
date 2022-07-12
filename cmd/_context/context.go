package _context

import (
	"context"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
)

type (
	Context interface {
		context.Context

		DBFile() dbfile.DBFile
		SetDBFile(dbFile dbfile.DBFile)
	}

	contextImpl struct {
		context.Context

		dbFile dbfile.DBFile
	}
)

func New(ctx context.Context) Context {
	return &contextImpl{
		Context: ctx,
	}
}

func (c *contextImpl) DBFile() dbfile.DBFile {
	return c.dbFile
}

func (c *contextImpl) SetDBFile(dbFile dbfile.DBFile) {
	c.dbFile = dbFile
}
