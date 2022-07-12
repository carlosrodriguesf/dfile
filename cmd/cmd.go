package cmd

import (
	"context"
	"github.com/carlosrodriguesf/dfile/cmd/_context"
	"github.com/carlosrodriguesf/dfile/cmd/path"
	"github.com/carlosrodriguesf/dfile/cmd/sum"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/pkg/logger"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

func getDefaultResourcePath() string {
	if p := os.Getenv("DFILE_RESOURCE_PATH"); p != "" {
		return p
	}
	return os.Getenv("HOME")
}

func startLogger(logFilePath string) (io.WriteCloser, error) {
	logWriter, err := logger.New(logFilePath)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logWriter)
	return logWriter, nil
}

func startDBFile(ctx _context.Context, dbFilePath string) error {
	dbFile := dbfile.New(dbFilePath, dbfile.Options{
		AutoPersist:      true,
		AutoPersistCount: 1000,
	})
	ctx.SetDBFile(dbFile)
	return dbFile.Load()
}

func Run() error {
	var (
		dbFilePath, logFilePath string
		logWriter               io.WriteCloser

		resourcePath = getDefaultResourcePath()
		ctx          = _context.New(context.Background())
		rootCmd      = cobra.Command{TraverseChildren: true}
	)

	rootCmd.AddCommand(path.Path(ctx))
	rootCmd.AddCommand(sum.Sum(ctx))

	rootCmd.Flags().StringVarP(&dbFilePath, "database", "d", resourcePath+"/dfile.db", "Database file")
	rootCmd.Flags().StringVarP(&logFilePath, "log-file", "l", resourcePath+"/dfile.log", "Log file")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) (err error) {
		if logWriter, err = startLogger(logFilePath); err != nil {
			log.Printf("error: %v", err)
			return err
		}
		if err = startDBFile(ctx, dbFilePath); err != nil {
			log.Printf("error: %v", err)
			return err
		}
		return nil
	}

	if logWriter != nil {
		defer logWriter.Close()
	}

	return rootCmd.Execute()
}
