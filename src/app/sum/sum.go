package sum

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/calculator"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
)

type (
	App interface {
		Generate(ctx context.Context) error
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
