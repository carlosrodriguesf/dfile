package dbm

import (
	"database/sql"
	"errors"
	"github.com/carlosrodriguesf/dfile/src/tool/dbm/migration"
	"github.com/carlosrodriguesf/dfile/src/tool/lh"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

func Open(file string) (*sql.DB, error) {
	exists, err := isDBExists(file)
	if err != nil {
		return nil, lh.LogError(err)
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, lh.LogError(err)
	}

	if !exists {
		log.Println("creating database")

		res, err := db.Exec(`
			CREATE TABLE db_version(
		    	version INTEGER NOT NULL
			);
			INSERT INTO db_version(version) VALUES (0);
		`)
		if err != nil {
			return nil, lh.LogError(err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return nil, lh.LogError(err)
		}
		if rowsAffected != 1 {
			return nil, errors.New("error when creating db")
		}
	}

	err = migration.Up(db)
	if err != nil {
		return nil, lh.LogError(err)
	}
	return db, nil
}

func isDBExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if err == os.ErrNotExist || strings.HasSuffix(err.Error(), "no such file or directory") {
		return false, nil
	}
	return false, lh.LogError(err)
}
