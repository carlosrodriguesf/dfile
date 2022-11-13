package cmd

import (
	"github.com/carlosrodrigues/dfile/cmd/path"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		TraverseChildren: true,
	}
	cmd.AddCommand(path.Command())
	return &cmd
}
