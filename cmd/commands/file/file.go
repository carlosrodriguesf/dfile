package file

import (
	"github.com/carlosrodriguesf/dfile/cmd/container"
	"github.com/spf13/cobra"
	"log"
)

func Command(container *container.Container) *cobra.Command {
	cmd := cobra.Command{
		Use:  "file",
		Args: cobra.ExactArgs(1),
	}
	cmd.AddCommand(commandGenerateSum(container))
	return &cmd
}

func commandGenerateSum(container *container.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "generate-sum",
		Short: "generate sum for all founded paths",
		Run: func(cmd *cobra.Command, args []string) {
			err := container.Service.File.GenerateChecksum(cmd.Context())
			if err != nil {
				log.Fatal(err)
			}
		},
	}
}
