package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/carlosrodriguesf/dfile/src/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/src/pkg/scanner"
	"log"
	"strings"
)

func (a *appImpl) Sync(ctx context.Context) error {
	dbFile := ctx.DBFile()

	for _, path := range dbFile.GetPathKeys() {
		err := a.syncPath(ctx, path, dbFile.GetPath(path))
		if err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}

	return nil
}

func (a *appImpl) syncPath(ctx context.Context, path string, entry dbfile.PathEntry) error {
	var (
		dbFile        = ctx.DBFile()
		filesFromDB   = make(map[string]bool)
		filesFromDisk = make(map[string]bool)
		files, err    = a.scanner.Scan(ctx, path, scanner.ScanOptions{
			AcceptExtensions: entry.AcceptExtensions,
			IgnoreFolders:    entry.IgnoreFolders,
		})
	)

	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	for _, file := range files {
		filesFromDisk[file] = true
	}

	for _, file := range dbFile.GetFileKeys() {
		if strings.HasPrefix(file, path) {
			filesFromDB[file] = true
		}
	}

	for file := range filesFromDisk {
		if !filesFromDB[file] {
			dbFile.SetFile(file, dbfile.FileEntry{
				Ready: false,
			})
		}
	}

	for file := range filesFromDB {
		if !filesFromDisk[file] {
			dbFile.DelFile(file)
		}
	}

	return dbFile.Persist()
}
