package path

import (
	"context"
	"github.com/carlosrodriguesf/dfile/pkg/model"
	"github.com/carlosrodriguesf/dfile/pkg/repository/file"
	"github.com/carlosrodriguesf/dfile/pkg/repository/path"
	"github.com/carlosrodriguesf/dfile/pkg/tool/array"
	"github.com/carlosrodriguesf/dfile/pkg/tool/scanner"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stack"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stacktrace"
	"log"
	"path/filepath"
)

type (
	Params struct {
		Scanner        scanner.Scanner
		PathRepository path.Repository
		FileRepository file.Repository
	}
	Service interface {
		Add(ctx context.Context, path model.Path) error
		Sync(ctx context.Context) (SyncResult, error)
	}
	service struct {
		scanner        scanner.Scanner
		pathRepository path.Repository
		fileRepository file.Repository
	}
)

func NewService(params Params) Service {
	return &service{
		scanner:        params.Scanner,
		pathRepository: params.PathRepository,
		fileRepository: params.FileRepository,
	}
}

func (s service) Add(ctx context.Context, path model.Path) error {
	absolutePath, err := filepath.Abs(path.Path)
	if err != nil {
		return stacktrace.WrapError(err)
	}

	path.Path = absolutePath
	err = s.pathRepository.Save(ctx, path)
	if err != nil {
		return stacktrace.WrapError(err)
	}

	_, err = s.Sync(ctx)
	if err != nil {
		return stacktrace.WrapError(err)
	}

	return err
}

func (s service) Sync(ctx context.Context) (SyncResult, error) {
	allPathsFromDB, err := s.pathRepository.ListAll(ctx)
	if err != nil {
		return SyncResult{}, stacktrace.WrapError(err)
	}

	allFilesFromDB, err := s.fileRepository.ListAll(ctx)
	if err != nil {
		return SyncResult{}, stacktrace.WrapError(err)
	}

	fileMapFromDB := array.ToBoolMapFunc(allFilesFromDB, func(v model.File) string {
		return v.Path
	})

	var result SyncResult
	for _, p := range allPathsFromDB {
		var pathResult SyncResult

		pathResult, err = s.syncPath(ctx, p, fileMapFromDB)
		if err != nil {
			return SyncResult{}, stacktrace.WrapError(err)
		}

		result.Added += pathResult.Added
		result.Removed += pathResult.Removed
	}
	return result, nil
}

func (s service) syncPath(ctx context.Context, path model.Path, filesFromDB map[string]bool) (SyncResult, error) {
	files, err := s.scanner.Scan(
		context.Background(), path.Path,
		scanner.WithValidExtensions(path.AcceptExtensions...),
		scanner.WithIgnoredNames(path.IgnoredFolderNames...),
	)
	if err != nil {
		return SyncResult{}, stacktrace.WrapError(err)
	}

	var (
		filesFromDisk = array.ToBoolMap(files)
		filesToRemove stack.Stack[string]
		filesToAdd    stack.Stack[model.File]
	)

	for f := range filesFromDisk {
		if !filesFromDB[f] {
			log.Printf("add '%s'\n", f)
			filesToAdd.Append(model.File{Path: f})
		}
	}

	for f, _ := range filesFromDB {
		if !filesFromDisk[f] {
			log.Printf("remove '%s'\n", f)
			filesToRemove.Append(f)
		}
	}

	if filesToAdd.Size() > 0 {
		err = s.fileRepository.SaveAll(ctx, filesToAdd.ToArray())
		if err != nil {
			return SyncResult{}, stacktrace.WrapError(err)
		}
	}

	if filesToRemove.Size() > 0 {
		err = s.fileRepository.RemoveAll(ctx, filesToRemove.ToArray())
		if err != nil {
			return SyncResult{}, stacktrace.WrapError(err)
		}
	}

	return SyncResult{
		Added:   filesToAdd.Size(),
		Removed: filesToRemove.Size(),
	}, nil
}
