package path

import (
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/repository/file"
	"github.com/carlosrodriguesf/dfile/src/repository/path"
	"github.com/carlosrodriguesf/dfile/src/tool/scanner"
)

type (
	Options struct {
		PathRep path.Repository
		FileRep file.Repository
		Scanner scanner.Scanner
	}

	App interface {
		Add(path model.Path) error
		Remove(path string) error
		List() ([]string, error)
	}

	appImpl struct {
		pathRep path.Repository
		fileRep file.Repository
		scanner scanner.Scanner
	}
)

func New(opts Options) App {
	return &appImpl{
		pathRep: opts.PathRep,
		fileRep: opts.FileRep,
		scanner: opts.Scanner,
	}
}
