package path

import (
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/app/path"
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/spf13/cobra"
)

func add(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "add [path]",
		Short: "add or remove path to sum.db",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			acceptExtensions, err := cmd.Flags().GetStringArray("extensions")
			if err != nil {
				return err
			}

			ignoreFolders, err := cmd.Flags().GetStringArray("ignore")
			if err != nil {
				return err
			}

			return app.Path().Add(ctx, args[0], path.AddConfig{
				AcceptExtensions: acceptExtensions,
				IgnoreFolders:    ignoreFolders,
			})
		},
	}
	cmd.Flags().StringArrayP("extensions", "e", []string{"jpg", "png", "mp4", "mov"}, "Extensões válidas.")
	cmd.Flags().StringArrayP("ignore", "i", []string{".debris", ".idea", ".git", ".gradle"}, "Extensões válidas.")

	return &cmd
}
