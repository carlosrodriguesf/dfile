package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"log"
	"path/filepath"
)

func (a appImpl) Remove(ctx context.Context, path string) error {
	dbFile := ctx.DBFile()

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	dbFile.DelPath(path)

	for _, file := range dbFile.GetFileKeys() {
		match, err := filepath.Match(path, file)
		if err != nil {
			log.Printf("error: %v", err)
			return err
		}
		if match {
			dbFile.DelFile(file)
		}
	}

	return dbFile.Persist()
}
