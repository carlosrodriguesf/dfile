package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/carlosrodriguesf/dfile/src/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/src/pkg/scanner"
)

type (
	AddConfig dbfile.PathEntry

	App interface {
		Add(ctx context.Context, path string, config AddConfig) error
		Remove(ctx context.Context, path string) error
		Sync(ctx context.Context) error
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
