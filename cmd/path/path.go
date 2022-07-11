package path

import (
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/spf13/cobra"
)

func Run(dbFile dbfile.DBFile) *cobra.Command {
	pathCmd := cobra.Command{
		Use:   "path [scan | remove] [path]",
		Short: "scan or remove path to sum.db",
		Args:  cobra.ExactArgs(1),
	}
	pathCmd.AddCommand(scan(dbFile))
	pathCmd.AddCommand(remove(dbFile))
	return &pathCmd
}
