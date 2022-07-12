package path

import (
	"github.com/carlosrodriguesf/dfile/cmd/_context"
	"github.com/spf13/cobra"
)

func Path(ctx _context.Context) *cobra.Command {
	pathCmd := cobra.Command{
		Use:   "path [scan | remove] [path]",
		Short: "scan or remove path to sum.db",
		Args:  cobra.ExactArgs(1),
	}
	pathCmd.AddCommand(scan(ctx))
	pathCmd.AddCommand(remove(ctx))
	return &pathCmd
}
