package path

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "path [add | remove] [path]",
		Short: "add or remove path",
		Args:  cobra.ExactArgs(1),
	}
	cmd.AddCommand(commandAdd())
	return &cmd
}

func commandAdd() *cobra.Command {
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

			fmt.Println("add command", acceptExtensions, ignoreFolders)

			return nil
		},
	}
	cmd.Flags().StringArrayP("extensions", "e", []string{"jpg", "png", "mp4", "mov"}, "Extensões válidas.")
	cmd.Flags().StringArrayP("ignore", "i", []string{".debris", ".idea", ".git", ".gradle"}, "Extensões válidas.")

	return &cmd
}
