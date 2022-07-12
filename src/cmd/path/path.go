package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
)

func Path(ctx context.Context) *cobra.Command {
	pathCmd := cobra.Command{
		Use:   "path [add | remove] [path]",
		Short: "add or remove path to sum.db",
		Args:  cobra.ExactArgs(1),
	}
	pathCmd.AddCommand(add(ctx))
	pathCmd.AddCommand(remove(ctx))
	pathCmd.AddCommand(sync(ctx))
	pathCmd.AddCommand(list(ctx))
	pathCmd.AddCommand(watch(ctx))
	return &pathCmd
}
