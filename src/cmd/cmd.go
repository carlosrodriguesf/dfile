package cmd

import (
	"github.com/carlosrodriguesf/dfile/src/cmd/db"
	"github.com/carlosrodriguesf/dfile/src/cmd/path"
	"github.com/carlosrodriguesf/dfile/src/cmd/sum"
	"github.com/carlosrodriguesf/dfile/src/cmd/watch"
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/carlosrodriguesf/dfile/src/tool/dbm"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"github.com/carlosrodriguesf/dfile/src/tool/logger"
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

func startLogger(logFilePath string, verbose bool) (io.WriteCloser, error) {
	logWriter, err := logger.New(logFilePath, verbose)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logWriter)
	return logWriter, nil
}

func startDBFile(ctx context.Context, dbFilePath string) error {
	db, err := dbm.Open(dbFilePath)
	if err != nil {
		return hlog.LogError(err)
	}
	ctx.Load(db)
	//dbFile := dbfile.New(dbFilePath, dbfile.Options{
	//	AutoPersist:      true,
	//	AutoPersistCount: 1000,
	//})
	//ctx.SetDBFile(dbFile)
	//return dbFile.Load()
	return nil
}

func Run() error {
	var (
		dbFilePath, logFilePath string
		verbose                 bool
		logWriter               io.WriteCloser

		resourcePath = getDefaultResourcePath()
		ctx          = context.New()
		rootCmd      = cobra.Command{TraverseChildren: true}
	)

	rootCmd.AddCommand(path.Path(ctx))
	rootCmd.AddCommand(sum.Sum(ctx))
	rootCmd.AddCommand(db.DB(ctx))
	rootCmd.AddCommand(watch.Watch(ctx))

	rootCmd.Flags().StringVarP(&dbFilePath, "database", "d", resourcePath+"/dfile.db", "Database file")
	rootCmd.Flags().StringVarP(&logFilePath, "log-file", "l", resourcePath+"/dfile.log", "Log file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show all logs in console output")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) (err error) {
		if logWriter, err = startLogger(logFilePath, verbose); err != nil {
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
