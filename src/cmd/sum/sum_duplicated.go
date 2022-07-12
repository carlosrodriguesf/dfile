package sum

import (
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
	"strings"
)

func duplicated(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "duplicated",
		Short: "get duplicated files",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			duplicated := app.Sum().Duplicated(ctx)
			template := "%s:\n\t%s\n\n"
			for hash, files := range duplicated {
				fmt.Printf(template, hash, strings.Join(files, "\n\t"))
			}
		},
	}
}
