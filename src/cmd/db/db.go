package db

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
)

func DB(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "db [rewrite]",
		Short: "rewrite database",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(rewrite(ctx))

	return &cmd
}
