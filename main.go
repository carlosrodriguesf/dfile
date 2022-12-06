package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/carlosrodriguesf/dfile/cmd"
	"github.com/carlosrodriguesf/dfile/cmd/container"
	"github.com/carlosrodriguesf/dfile/pkg/tool/db"
	"github.com/carlosrodriguesf/dfile/pkg/tool/scanner"
	"io"
	"log"
	"os"
	"strconv"
	"syscall"
)

func getLogWriter(outputFile string) (io.WriteCloser, error) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0700)
	if errors.Is(err, os.ErrNotExist) || errors.Is(err, syscall.ENOENT) {
		file, err = os.Create(outputFile)
	}
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	log.SetPrefix(strconv.Itoa(os.Getpid()) + " ")

	fileWriter, err := getLogWriter("dfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer fileWriter.Close()

	log.SetOutput(io.MultiWriter(fileWriter, os.Stdout))
	dbConn, err := db.Open("dfile.db")
	if err != nil {
		log.Fatal(err)
	}
	appContainer := container.New(dbConn)
	if err := cmd.Command(appContainer).Execute(); err != nil {
		log.Fatal(err)
	}
}

func r() {
	fileWriter, err := getLogWriter("dfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer fileWriter.Close()

	log.SetOutput(io.MultiWriter(fileWriter, os.Stdout))
	//dbConn, err := db.Open("dfile.db")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//appContainer := container.New(dbConn)
	//if err := cmd.Command(appContainer).Execute(); err != nil {
	//	log.Fatal(err)
	//}

	log.SetPrefix(strconv.Itoa(os.Getpid()) + " ")
	scn := scanner.New()

	opts := []scanner.Option{
		scanner.WithIgnoredNames("Workspace", "node_modules", ".local", ".libs", "model", ".npm", ".cache", ".idea", ".config"),
		scanner.WithValidExtensions("jpg", "png", "go"),
	}
	files, err := scn.Scan(context.Background(), "/", opts...)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("%v %+v \n", err, files)
}
