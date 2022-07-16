package app

import (
	"github.com/carlosrodriguesf/dfile/src/app/path"
	"github.com/carlosrodriguesf/dfile/src/app/sum"
	"github.com/carlosrodriguesf/dfile/src/app/watch"
	"github.com/carlosrodriguesf/dfile/src/repository"
	"github.com/carlosrodriguesf/dfile/src/tool/calculator"
	"github.com/carlosrodriguesf/dfile/src/tool/scanner"
)

type (
	Container interface {
		Path() path.App
		Sum() sum.App
		Watch() watch.App
	}
	container struct {
		repository repository.Container

		path  path.App
		sum   sum.App
		watch watch.App
	}
)

func NewContainer(repositoryContainer repository.Container) Container {
	return &container{
		repository: repositoryContainer,
	}
}

func (c *container) Path() path.App {
	if c.path == nil {
		c.path = path.New(scanner.New(), c.Sum())
	}
	return c.path
}

func (c *container) Sum() sum.App {
	if c.sum == nil {
		c.sum = sum.New(
			calculator.New(),
		)
	}
	return c.sum
}

func (c *container) Watch() watch.App {
	if c.watch == nil {
		c.watch = watch.New(c.Sum())
	}
	return c.watch
}
