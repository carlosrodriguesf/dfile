package container

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/cmd/container/repository"
	"github.com/carlosrodriguesf/dfile/cmd/container/service"
	"github.com/carlosrodriguesf/dfile/pkg/tool/scanner"
	"github.com/carlosrodriguesf/dfile/pkg/tool/sumcalc"
)

type Container struct {
	Service    *service.Container
	Repository *repository.Container
}

func New(db *sql.DB) *Container {
	repositoryContainer := repository.NewContainer(repository.Params{DB: db})
	serviceContainer := service.NewContainer(service.Params{
		Scanner:    scanner.New(),
		SumCalc:    sumcalc.New(),
		Repository: repositoryContainer,
	})
	return &Container{
		Repository: repositoryContainer,
		Service:    serviceContainer,
	}
}
