package path

import (
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
)

func watch(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "generate sum for scanned paths",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Path().Watch(ctx)
		},
	}
}
