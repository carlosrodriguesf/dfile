package migration

import (
	"database/sql"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
)

type migrateFunc func(db *sql.DB) error

var migrations = []migrateFunc{
	func(db *sql.DB) error {
		_, err := db.Exec(`
			CREATE TABLE paths(
			    path TEXT NOT NULL PRIMARY KEY,
			    enabled INTEGER NOT NULL DEFAULT 1,
			    ignore_folders TEXT,
			    accept_extensions TEXT
			);

			CREATE TABLE files(
			    path TEXT NOT NULL PRIMARY KEY,
			    checksum TEXT,
			    error TEXT
			);
		`)
		return err
	},
}

func Up(db *sql.DB) error {
	version, err := getVersion(db)
	if err != nil {
		return hlog.LogError(err)
	}

	lastVersion := len(migrations)
	for i := version; i < lastVersion; i++ {
		err = migrations[i](db)
		if err != nil {
			return hlog.LogError(err)
		}
	}

	err = setVersion(db, lastVersion)
	if err != nil {
		return hlog.LogError(err)
	}

	return nil
}

func setVersion(db *sql.DB, version int) error {
	_, err := db.Exec("UPDATE db_version SET version = $1", version)
	if err != nil {
		return hlog.LogError(err)
	}
	return nil
}

func getVersion(db *sql.DB) (version int, err error) {
	res, err := db.Query("SELECT version FROM db_version")
	if err != nil {
		return 0, hlog.LogError(err)
	}
	defer hlog.LogClose(res)

	if !res.Next() {
		return 0, nil
	}
	err = res.Scan(&version)
	if err != nil {
		return 0, hlog.LogError(err)
	}
	return version, nil
}
