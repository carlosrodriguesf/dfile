package path

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/spf13/cobra"
)

func remove(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "remove [path]",
		Short: "remove path from sum.db",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ctx.App().Path().Remove(args[0])
		},
	}
	return &cmd
}
