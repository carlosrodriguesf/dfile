package watch

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/spf13/cobra"
)

func Watch(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "watch for file changes",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return ctx.App().Watch().Watch()
		},
	}
}
