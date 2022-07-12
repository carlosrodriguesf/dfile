package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/carlosrodriguesf/dfile/src/pkg/scanner"
	"path/filepath"
)

func (a *appImpl) Add(ctx context.Context, path string, opts ...AddOptions) error {
	var (
		dbFile = ctx.DBFile()
		opt    scanner.ScanOptions
	)

	if len(opts) > 0 {
		opt = scanner.ScanOptions(opts[0])
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	files, err := a.scanner.Scan(ctx, path, opt)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !dbFile.Has(file) {
			dbFile.CreateEntry(file)
		}
	}

	err = dbFile.Persist()
	if err != nil {
		return err
	}

	return nil
}
