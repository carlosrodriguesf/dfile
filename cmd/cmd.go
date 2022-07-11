package cmd

import (
	"github.com/carlosrodriguesf/dfile/cmd/path"
	"github.com/carlosrodriguesf/dfile/cmd/sum"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/spf13/cobra"
	"os"
)

func Run(dbFile dbfile.DBFile) error {
	resourcePath := os.Getenv("DFILE_RESOURCE_PATH")
	if resourcePath == "" {
		resourcePath = os.Getenv("HOME")
	}

	rootCmd := cobra.Command{}
	rootCmd.AddCommand(path.Run(dbFile))
	rootCmd.AddCommand(sum.Sum(dbFile))
	return rootCmd.Execute()
}
