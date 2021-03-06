package sum

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/spf13/cobra"
)

func generate(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "generate sum for scanner paths",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ctx.App().Sum().Generate()
		},
	}
}
