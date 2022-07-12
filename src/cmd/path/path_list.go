package path

import (
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
)

func list(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all paths in db",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			for _, path := range app.Path().List(ctx) {
				fmt.Println(path)
			}
		},
	}
}
