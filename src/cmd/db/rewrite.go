package db

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/carlosrodriguesf/dfile/src/tool/dbfile"
	"github.com/spf13/cobra"
	"log"
)

func rewrite(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:       "rewrite",
		Short:     "rewrite database",
		ValidArgs: []string{"indented", "not-indented"},
		RunE: func(cmd *cobra.Command, args []string) error {
			indented, err := cmd.Flags().GetBool("indented")
			if err != nil {
				log.Printf("error: %v", err)
				return err
			}

			return ctx.DBFile().Persist(dbfile.PersistOptions{
				Indented: indented,
			})
		},
	}

	cmd.Flags().BoolP("indented", "i", false, "save database indentend")

	return &cmd
}
