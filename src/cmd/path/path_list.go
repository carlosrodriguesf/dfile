package path

import (
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"github.com/spf13/cobra"
)

func list(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all paths in db",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := ctx.App().Path().List()
			if err != nil {
				return hlog.LogError(err)
			}
			for _, path := range paths {
				fmt.Println(path)
			}
			return nil
		},
	}
}
