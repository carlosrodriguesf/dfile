package path

import (
	"github.com/carlosrodriguesf/dfile/cmd/container"
	"github.com/carlosrodriguesf/dfile/pkg/model"
	"github.com/spf13/cobra"
	"log"
)

func Command(container *container.Container) *cobra.Command {
	cmd := cobra.Command{
		Use:   "path [add | remove] [path]",
		Short: "add or remove path",
		Args:  cobra.ExactArgs(1),
	}
	cmd.AddCommand(commandAdd(container))
	cmd.AddCommand(commandSync(container))
	return &cmd
}

func commandAdd(container *container.Container) *cobra.Command {
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

			err = container.Service.Path.Add(cmd.Context(), model.Path{
				Path:               args[0],
				Enabled:            true,
				AcceptExtensions:   acceptExtensions,
				IgnoredFolderNames: ignoreFolders,
			})
			if err != nil {
				log.Println(err)
			}
			return nil
		},
	}
	cmd.Flags().StringArrayP("extensions", "e", []string{"jpg", "png", "mp4", "mov"}, "Extensões válidas.")
	cmd.Flags().StringArrayP("ignore", "i", []string{".debris", ".idea", ".git", ".gradle"}, "Extensões válidas.")

	return &cmd
}

func commandSync(container *container.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "sync all registered paths",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := container.Service.Path.Sync(cmd.Context())
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%d files added\n", result.Added)
			log.Printf("%d files removed\n", result.Removed)
		},
	}
}
