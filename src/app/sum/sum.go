package sum

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/calculator"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
)

type (
	App interface {
		Generate(ctx context.Context) error
		Duplicated(ctx context.Context) map[string][]string
	}

	appImpl struct {
		calculator calculator.Calculator
	}
)

func New(calculator calculator.Calculator) App {
	return &appImpl{
		calculator: calculator,
	}
}
