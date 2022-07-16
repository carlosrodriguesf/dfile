package path

import (
	"database/sql"
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/dbm"
	"github.com/carlosrodriguesf/dfile/src/tool/lh"
)

type (
	Repository interface {
		Save(path model.Path) error
		Get(path string) (model.Path, error)
		GetPaths() ([]string, error)
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

func (r *repository) Save(path model.Path) error {
	query := `INSERT INTO paths(path, enabled, ignore_folders, accept_extensions) VALUES (?,?,?,?)`

	res, err := r.db.Exec(
		query,
		path.Path,
		path.Enabled,
		dbm.JSONArray(path.IgnoreFolders),
		dbm.JSONArray(path.AcceptExtensions),
	)
	fmt.Println("exec: ", query, err)
	rr, _ := res.RowsAffected()
	fmt.Println(rr)
	if err != nil {
		return lh.LogError(err)
	}
	return nil
}

func (r *repository) Remove(path string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Get(path string) (model.Path, error) {
	var data model.Path

	query := `SELECT path, enabled, ignore_folders, accept_extensions FROM paths WHERE path = ?`
	res, err := r.db.Query(query, path)
	if err != nil {
		return data, lh.LogError(err)
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
		return data, lh.LogError(err)
	}
	return data, nil
}

func (r *repository) GetPaths() ([]string, error) {
	//TODO implement me
	panic("implement me")
}
