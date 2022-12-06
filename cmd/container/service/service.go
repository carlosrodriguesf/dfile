package service

import (
	"github.com/carlosrodriguesf/dfile/cmd/container/repository"
	"github.com/carlosrodriguesf/dfile/pkg/service/file"
	"github.com/carlosrodriguesf/dfile/pkg/service/path"
	"github.com/carlosrodriguesf/dfile/pkg/tool/scanner"
	"github.com/carlosrodriguesf/dfile/pkg/tool/sumcalc"
)

type (
	Params struct {
		Scanner    scanner.Scanner
		SumCalc    sumcalc.Calculator
		Repository *repository.Container
	}
	Container struct {
		Path path.Service
		File file.Service
	}
)

func NewContainer(params Params) *Container {
	return &Container{
		Path: path.NewService(path.Params{
			Scanner:        params.Scanner,
			PathRepository: params.Repository.Path,
			FileRepository: params.Repository.File,
		}),
		File: file.NewService(file.Params{
			SumCalc:        params.SumCalc,
			FileRepository: params.Repository.File,
		}),
	}
}
