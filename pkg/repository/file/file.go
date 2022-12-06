package file

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carlosrodriguesf/dfile/pkg/model"
	"github.com/carlosrodriguesf/dfile/pkg/tool/array"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stacktrace"
	"strings"
)

type (
	Repository interface {
		ListAll(ctx context.Context) ([]model.File, error)
		SaveAll(ctx context.Context, file []model.File) error
		Save(ctx context.Context, file model.File) error
		RemoveAll(ctx context.Context, paths []string) error
	}
	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) ListAll(ctx context.Context) ([]model.File, error) {
	query := `SELECT path, checksum, error FROM files`
	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, stacktrace.WrapError(err)
	}
	defer res.Close()

	files := make([]model.File, 0)
	for res.Next() {
		var file model.File
		err = res.Scan(
			&file.Path,
			&file.Checksum,
			&file.Error,
		)
		if err != nil {
			return nil, stacktrace.WrapError(err)
		}
		files = append(files, file)
	}

	if err != nil {
		return nil, stacktrace.WrapError(err)
	}
	return files, nil
}

func (r repository) SaveAll(ctx context.Context, files []model.File) error {
	chunks := array.Chunk(files, 100)
	for _, chunk := range chunks {
		err := r.saveChunk(ctx, chunk)
		if err != nil {
			return stacktrace.WrapError(err)
		}
	}
	return nil
}

func (r repository) Save(ctx context.Context, file model.File) error {
	query := `INSERT OR REPLACE INTO files(path, checksum, error) VALUES (?,?,?)`
	_, err := r.db.ExecContext(
		ctx, query,
		file.Path,
		file.Checksum,
		file.Error,
	)
	if err != nil {
		return stacktrace.WrapError(err)
	}
	return nil
}

func (r repository) RemoveAll(ctx context.Context, path []string) error {
	for _, chunk := range array.Chunk(path, 300) {
		err := r.removeChunk(ctx, chunk)
		if err != nil {
			return stacktrace.WrapError(err)
		}
	}
	return nil
}

func (r repository) saveChunk(ctx context.Context, files []model.File) error {
	valuesQuery := make([]string, len(files))
	valuesParam := make([]interface{}, len(files)*3)
	for i := 0; i < len(files); i++ {
		valuesQuery[i] = `(?,?,?)`

		j := i * 3
		valuesParam[j] = files[i].Path
		valuesParam[j+1] = files[i].Checksum
		valuesParam[j+2] = files[i].Error
	}

	query := fmt.Sprintf("INSERT INTO files(path, checksum, error) VALUES %s", strings.Join(valuesQuery, ","))
	_, err := r.db.ExecContext(ctx, query, valuesParam...)
	if err != nil {
		return stacktrace.WrapError(err)
	}
	return nil
}

func (r repository) removeChunk(ctx context.Context, paths []string) error {
	var (
		params = make([]string, len(paths))
		values = make([]any, len(paths))
	)
	for i := range paths {
		params[i] = "?"
		values[i] = paths[i]
	}

	query := fmt.Sprintf("DELETE FROM files WHERE path IN (%s)", strings.Join(params, ","))
	res, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return stacktrace.WrapError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return stacktrace.WrapError(err)
	}
	if rowsAffected == 0 {
		return stacktrace.WrapError(sql.ErrNoRows)
	}
	return nil
}
