package sum

import (
	"github.com/carlosrodriguesf/dfile/src/tool/calculator"
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"sync"
)

type (
	App interface {
		Generate(ctx context.Context) error
		GenerateFileSum(ctx context.Context, file string, persist bool)
		Duplicated(ctx context.Context) map[string][]string
	}

	appImpl struct {
		calculator calculator.Calculator
		mutex      *sync.Mutex
	}
)

func New(calculator calculator.Calculator) App {
	return &appImpl{
		calculator: calculator,
	}
}
