package path

import (
	"github.com/carlosrodriguesf/dfile/cmd/_context"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

func remove(ctx _context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "remove [path]",
		Short: "remove path from sum.db",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dbFile := ctx.DBFile()

			path, err := filepath.Abs(args[0])
			if err != nil {
				return err
			}

			for _, file := range dbFile.Keys() {
				match, err := filepath.Match(path, file)
				if err != nil {
					log.Printf("error: %v", err)
					return err
				}
				if match {
					dbFile.Del(file)
				}
			}

			return dbFile.Persist()
		},
	}
	return &cmd
}
