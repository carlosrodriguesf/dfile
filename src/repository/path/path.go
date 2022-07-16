package path

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/dbm"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
)

type (
	Repository interface {
		All() ([]model.Path, error)
		Get(path string) (model.Path, error)
		Has(path string) (bool, error)
		Save(path model.Path) error
		Remove(path string) error
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

func (r *repository) All() ([]model.Path, error) {
	query := `SELECT path, enabled, ignore_folders, accept_extensions FROM paths`
	res, err := r.db.Query(query)
	if err != nil {
		return nil, hlog.LogError(err)
	}
	defer res.Close()

	paths := make([]model.Path, 0)
	if res.Next() {
		var path model.Path
		err = res.Scan(
			&path.Path,
			&path.Enabled,
			dbm.JSONArray(&path.IgnoreFolders),
			dbm.JSONArray(&path.AcceptExtensions),
		)
		if err != nil {
			return nil, hlog.LogError(err)
		}
		paths = append(paths, path)
	}

	if err != nil {
		return nil, hlog.LogError(err)
	}

	return paths, nil
}

func (r *repository) Get(path string) (model.Path, error) {
	var data model.Path

	query := `SELECT path, enabled, ignore_folders, accept_extensions FROM paths WHERE path = ?`
	res, err := r.db.Query(query, path)
	if err != nil {
		return data, hlog.LogError(err)
	}
	if !res.Next() {
		return data, sql.ErrNoRows
	}

	err = res.Scan(
		&data.Path,
		&data.Enabled,
		dbm.JSONArray(&data.IgnoreFolders),
		dbm.JSONArray(&data.AcceptExtensions),
	)
	if err != nil {
		return data, hlog.LogError(err)
	}
	return data, nil
}

func (r *repository) Has(path string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM paths WHERE path = ?)`
	res, err := r.db.Query(query, path)
	if err != nil {
		return false, hlog.LogError(err)
	}
	defer res.Close()

	res.Next()

	var exists bool
	err = res.Scan(&exists)
	if err != nil {
		return false, hlog.LogError(err)
	}
	return exists, nil
}

func (r *repository) Save(path model.Path) error {
	query := `INSERT INTO paths(path, enabled, ignore_folders, accept_extensions) VALUES (?,?,?,?)`
	_, err := r.db.Exec(
		query,
		path.Path,
		path.Enabled,
		dbm.JSONArray(path.IgnoreFolders),
		dbm.JSONArray(path.AcceptExtensions),
	)
	if err != nil {
		return hlog.LogError(err)
	}
	return nil
}

func (r *repository) Remove(path string) error {
	query := `DELETE FROM paths WHERE path = ?`
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
