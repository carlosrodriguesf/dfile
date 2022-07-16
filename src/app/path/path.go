package path

import (
	"github.com/carlosrodriguesf/dfile/src/app/sum"
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
		List(ctx context.Context) []string
	}

	appImpl struct {
		scanner scanner.Scanner
		sum     sum.App
	}
)

func New(scanner scanner.Scanner, sum sum.App) App {
	return &appImpl{
		scanner: scanner,
		sum:     sum,
	}
}
