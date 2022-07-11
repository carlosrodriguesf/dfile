package cmd

import (
	"github.com/carlosrodriguesf/dfile/cmd/path"
	"github.com/carlosrodriguesf/dfile/cmd/sum"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/spf13/cobra"
)

func Run(dbFile dbfile.DBFile) error {
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(path.Path(dbFile))
	rootCmd.AddCommand(sum.Sum(dbFile))
	return rootCmd.Execute()
}
