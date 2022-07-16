package app

import (
	"github.com/carlosrodriguesf/dfile/src/app/path"
	"github.com/carlosrodriguesf/dfile/src/app/sum"
	"github.com/carlosrodriguesf/dfile/src/app/watch"
	"github.com/carlosrodriguesf/dfile/src/pkg/calculator"
	"github.com/carlosrodriguesf/dfile/src/pkg/scanner"
)

var container = struct {
	path  path.App
	sum   sum.App
	watch watch.App
}{}

func Path() path.App {
	if container.path == nil {
		container.path = path.New(scanner.New(), Sum())
	}
	return container.path
}

func Sum() sum.App {
	if container.sum == nil {
		container.sum = sum.New(
			calculator.New(),
		)
	}
	return container.sum
}

func Watch() watch.App {
	if container.watch == nil {
		container.watch = watch.New(Sum())
	}
	return container.watch
}
