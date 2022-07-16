package path

import (
	"context"
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"github.com/carlosrodriguesf/dfile/src/tool/scanner"
	"log"
	"strings"
)

func (a *appImpl) Sync() error {
	allPaths, err := a.pathRep.All()
	if err != nil {
		return hlog.LogError(err)
	}

	allFiles, err := a.fileRep.All()
	if err != nil {
		return hlog.LogError(err)
	}

	for _, path := range allPaths {
		err = a.syncPath(path, allFiles)
		if err != nil {
			return hlog.LogError(err)
		}
	}

	return nil
}

func (a *appImpl) syncPath(path model.Path, allFiles []model.File) error {
	var (
		filesFromDB   = make(map[string]bool)
		filesFromDisk = make(map[string]bool)
		files, err    = a.scanner.Scan(context.Background(), path.Path, scanner.ScanOptions{
			AcceptExtensions: path.AcceptExtensions,
			IgnoreFolders:    path.IgnoreFolders,
		})
	)

	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	for _, file := range files {
		filesFromDisk[file] = true
	}

	for _, file := range allFiles {
		if strings.HasPrefix(file.Path, path.Path) {
			filesFromDB[file.Path] = true
		}
	}

	var (
		filesToRemove = make([]string, 0)
		filesToAdd    = make([]model.File, 0)
	)

	for file := range filesFromDisk {
		if !filesFromDB[file] {
			filesToAdd = append(filesToAdd, model.File{Path: file})
		}
	}

	for file := range filesFromDB {
		if !filesFromDisk[file] {
			filesToRemove = append(filesToRemove, file)
		}
	}

	// TODO: a.fileRep.SaveAll(filesToAdd)
	// TODO: a.fileRep.DeleteAll(filesToDelete)

	return nil
}
