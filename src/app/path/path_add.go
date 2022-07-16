package path

import (
	"github.com/carlosrodriguesf/dfile/src/tool/apperrors"
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/carlosrodriguesf/dfile/src/tool/dbfile"
	"github.com/carlosrodriguesf/dfile/src/tool/scanner"
	"path/filepath"
)

func (a *appImpl) Add(ctx context.Context, path string, config AddConfig) error {
	dbFile := ctx.DBFile()

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if dbFile.HasPath(path) {
		return apperrors.ErrPathAlreadyExists
	}

	dbFile.SetPath(path, dbfile.PathEntry(config))

	files, err := a.scanner.Scan(ctx, path, scanner.ScanOptions{
		AcceptExtensions: config.AcceptExtensions,
		IgnoreFolders:    config.IgnoreFolders,
	})
	if err != nil {
		return err
	}

	for _, file := range files {
		if !dbFile.HasFile(file) {
			dbFile.SetFile(file, dbfile.FileEntry{
				Ready: false,
			})
		}
	}

	err = dbFile.Persist()
	if err != nil {
		return err
	}

	return nil
}
