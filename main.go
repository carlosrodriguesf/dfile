package main

import (
	"github.com/carlosrodriguesf/dfile/cmd"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/pkg/logger"
	"log"
	"os"
)

func main() {
	resourcePath := os.Getenv("DFILE_RESOURCE_PATH")
	if resourcePath == "" {
		resourcePath = os.Getenv("HOME")
	}

	customLogger, err := logger.New(resourcePath + "/dfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer customLogger.Close()

	log.SetOutput(customLogger)

	dbFile := dbfile.New(resourcePath+"/dfile.db", dbfile.Options{
		AutoPersist:      true,
		AutoPersistCount: 1000,
	})
	if err := dbFile.Load(); err != nil {
		log.Fatal(err)
	}

	err = cmd.Run(dbFile)
	if err != nil {
		log.Fatal(err)
	}
}
