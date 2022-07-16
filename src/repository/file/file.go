package file

import (
	"database/sql"
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"strings"
)

type (
	Repository interface {
		All() ([]model.File, error)
		Get(file string) (model.File, error)
		Save(file model.File) error
		SaveAll(files []model.File) error
		Remove(file string) error
		ExistingFiles() (map[string]bool, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) All() ([]model.File, error) {
	query := `SELECT path, checksum, error FROM files`
	res, err := r.db.Query(query)
	if err != nil {
		return nil, hlog.LogError(err)
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
			return nil, hlog.LogError(err)
		}
		files = append(files, file)
	}

	if err != nil {
		return nil, hlog.LogError(err)
	}
	return files, nil
}

func (r *repository) Get(path string) (model.File, error) {
	var file model.File

	query := `SELECT path, checksum, error FROM files WHERE path = ?`
	res, err := r.db.Query(query, path)
	if err != nil {
		return file, hlog.LogError(err)
	}
	defer res.Close()

	if !res.Next() {
		return file, sql.ErrNoRows
	}

	err = res.Scan(
		&file.Path,
		&file.Checksum,
		&file.Error,
	)
	if err != nil {
		return file, hlog.LogError(err)
	}
	return file, nil
}

func (r *repository) Save(path model.File) error {
	query := `INSERT OR REPLACE INTO files(path, checksum, error) VALUES (?,?,?)`
	_, err := r.db.Exec(
		query,
		path.Path,
		path.Checksum,
		path.Error,
	)
	if err != nil {
		return hlog.LogError(err)
	}
	return nil
}

func (r *repository) SaveAll(files []model.File) error {
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
	_, err := r.db.Exec(query, valuesParam...)
	if err != nil {
		return hlog.LogError(err)
	}
	return nil
}

func (r *repository) Remove(path string) error {
	query := `DELETE FROM files WHERE path = ?`
	res, err := r.db.Exec(query, path)
	if err != nil {
		return hlog.LogError(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return hlog.LogError(err)
	}
	if rowsAffected == 0 {
		return hlog.LogError(sql.ErrNoRows)
	}
	return nil
}

func (r *repository) ExistingFiles() (map[string]bool, error) {
	query := `SELECT path FROM files`
	res, err := r.db.Query(query)
	if err != nil {
		return nil, hlog.LogError(err)
	}
	defer res.Close()

	fileMap := make(map[string]bool)
	for res.Next() {
		var file string

		err = res.Scan(&file)
		if err != nil {
			return nil, hlog.LogError(err)
		}

		fileMap[file] = true
	}

	return fileMap, nil
}
