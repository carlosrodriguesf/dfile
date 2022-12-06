package cmd

import (
	"github.com/carlosrodriguesf/dfile/cmd/commands/file"
	"github.com/carlosrodriguesf/dfile/cmd/commands/path"
	"github.com/carlosrodriguesf/dfile/cmd/container"
	"github.com/spf13/cobra"
)

func Command(container *container.Container) *cobra.Command {
	cmd := cobra.Command{
		TraverseChildren: true,
	}
	cmd.AddCommand(path.Command(container))
	cmd.AddCommand(file.Command(container))
	return &cmd
}
