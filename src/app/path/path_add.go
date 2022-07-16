package path

import (
	"context"
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/apperrors"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"github.com/carlosrodriguesf/dfile/src/tool/scanner"
	"path/filepath"
)

func (a *appImpl) Add(path model.Path) error {
	var err error

	path.Path, err = filepath.Abs(path.Path)
	if err != nil {
		return err
	}

	exists, err := a.pathRep.Has(path.Path)
	if err != nil {
		return hlog.LogError(err)
	}
	if exists {
		return apperrors.ErrPathAlreadyExists
	}

	err = a.pathRep.Save(path)
	if err != nil {
		return hlog.LogError(err)
	}

	return a.scan(path)
}

func (a *appImpl) scan(path model.Path) error {
	files, err := a.scanner.Scan(context.Background(), path.Path, scanner.ScanOptions{
		AcceptExtensions: path.AcceptExtensions,
		IgnoreFolders:    path.IgnoreFolders,
	})
	if err != nil {
		return hlog.LogError(err)
	}

	existingFiles, err := a.fileRep.ExistingFiles()
	if err != nil {
		return hlog.LogError(err)
	}

	filesToSave := make([]model.File, 0)
	for _, file := range files {
		if !existingFiles[file] {
			filesToSave = append(filesToSave, model.File{
				Path: file,
			})
		}
	}

	err = a.fileRep.SaveAll(filesToSave)
	if err != nil {
		return hlog.LogError(err)
	}

	return nil
}
