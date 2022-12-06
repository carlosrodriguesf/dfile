package path

import (
	"context"
	"database/sql"
	"github.com/carlosrodriguesf/dfile/pkg/model"
	"github.com/carlosrodriguesf/dfile/pkg/tool/db"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stacktrace"
)

type (
	Repository interface {
		Save(ctx context.Context, path model.Path) error
		ListAll(ctx context.Context) ([]model.Path, error)
	}
	repositoryImpl struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) Save(ctx context.Context, path model.Path) error {
	query := `INSERT INTO paths(path, enabled, ignore_folders, accept_extensions) VALUES (?,?,?,?)`
	_, err := r.db.ExecContext(
		ctx,
		query,
		path.Path,
		path.Enabled,
		db.JSONArray(path.IgnoredFolderNames),
		db.JSONArray(path.AcceptExtensions),
	)
	if err != nil {
		return stacktrace.WrapError(err)
	}
	return nil
}

func (r *repositoryImpl) ListAll(ctx context.Context) ([]model.Path, error) {
	query := `SELECT path, enabled, ignore_folders, accept_extensions FROM paths`
	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, stacktrace.WrapError(err)
	}
	defer res.Close()

	paths := make([]model.Path, 0)
	if res.Next() {
		var path model.Path
		err = res.Scan(
			&path.Path,
			&path.Enabled,
			db.JSONArray(&path.IgnoredFolderNames),
			db.JSONArray(&path.AcceptExtensions),
		)
		if err != nil {
			return nil, stacktrace.WrapError(err)
		}
		paths = append(paths, path)
	}

	if err != nil {
		return nil, stacktrace.WrapError(err)
	}

	return paths, nil
}
