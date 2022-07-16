package repository

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/repository/path"
)

var container = struct {
	path path.Repository
}{}

func Path(db *sql.DB) path.Repository {
	return path.NewRepository(db)
}
