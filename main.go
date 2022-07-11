package main

import (
	"github.com/carlosrodriguesf/dfile/cmd"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/pkg/logger"
	"log"
	"os"
)

func main() {
	customLogger, err := logger.New("output.log")
	if err != nil {
		log.Fatal(err)
	}
	defer customLogger.Close()

	log.SetOutput(customLogger)

	dbFilePath := os.Getenv("DFILE_DB_PATH")
	if dbFilePath == "" {
		dbFilePath = os.Getenv("HOME") + "/dfile.db"
	}

	dbFile := dbfile.New(dbFilePath)
	if err := dbFile.Load(); err != nil {
		log.Fatal(err)
	}

	err = cmd.Run(dbFile)
	if err != nil {
		log.Fatal(err)
	}
}
