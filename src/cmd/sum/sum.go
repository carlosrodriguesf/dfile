package sum

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
)

func Sum(ctx context.Context) *cobra.Command {
	pathCmd := cobra.Command{
		Use:   "sum [generate]",
		Short: "generate sum for scanned paths",
		Args:  cobra.ExactArgs(1),
	}
	pathCmd.AddCommand(generate(ctx))
	pathCmd.AddCommand(duplicated(ctx))
	return &pathCmd
}
