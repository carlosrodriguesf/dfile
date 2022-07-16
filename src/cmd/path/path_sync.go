package path

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/spf13/cobra"
)

func sync(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "sync path",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			//return ctx.App().Path().Sync(ctx)
			return nil
		},
	}
}
