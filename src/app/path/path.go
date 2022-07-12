package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/carlosrodriguesf/dfile/src/pkg/scanner"
)

type (
	AddOptions scanner.ScanOptions

	App interface {
		Add(ctx context.Context, path string, opts ...AddOptions) error
		Remove(ctx context.Context, path string) error
	}

	appImpl struct {
		scanner scanner.Scanner
	}
)

func New(scanner scanner.Scanner) App {
	return &appImpl{
		scanner: scanner,
	}
}
