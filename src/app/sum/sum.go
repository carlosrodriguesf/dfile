package sum

import (
	"github.com/carlosrodriguesf/dfile/src/repository/file"
	"github.com/carlosrodriguesf/dfile/src/tool/calculator"
	"sync"
)

type (
	Options struct {
		FileRep    file.Repository
		Calculator calculator.Calculator
	}

	App interface {
		Generate() error
		GenerateFileSum(file string)
		Duplicated() (map[string][]string, error)
	}

	appImpl struct {
		fileRep    file.Repository
		calculator calculator.Calculator
		mutex      *sync.Mutex
	}
)

func New(opts Options) App {
	return &appImpl{
		fileRep:    opts.FileRep,
		calculator: opts.Calculator,
	}
}
