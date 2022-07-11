package path

import (
	"context"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/pkg/scanner"
	"github.com/spf13/cobra"
	"path/filepath"
)

func scan(dbFile dbfile.DBFile) *cobra.Command {
	cmd := cobra.Command{
		Use:   "scan [path]",
		Short: "scan or remove path to sum.db",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := filepath.Abs(args[0])
			if err != nil {
				return err
			}

			acceptExtensions, err := cmd.Flags().GetStringArray("extensions")
			if err != nil {
				return err
			}

			ignoreFolders, err := cmd.Flags().GetStringArray("ignore")
			if err != nil {
				return err
			}

			s := scanner.New(scanner.Options{
				AcceptExtensions: acceptExtensions,
				IgnoreFolders:    ignoreFolders,
			})

			files, err := s.Scan(context.Background(), path)
			if err != nil {
				return err
			}

			for _, file := range files {
				if !dbFile.Has(file) {
					dbFile.CreateEntry(file)
				}
			}

			err = dbFile.Persist()
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().StringArrayP("extensions", "e", []string{"jpg", "png", "mp4", "mov"}, "Extensões válidas.")
	cmd.Flags().StringArrayP("ignore", "i", []string{"jpg", "png", "mp4", "mov"}, "Extensões válidas.")

	return &cmd
}
