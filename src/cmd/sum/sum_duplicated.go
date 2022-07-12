package sum

import (
	"encoding/json"
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/app"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

func duplicated(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "duplicated",
		Short: "get duplicated files",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			outputFormat, err := cmd.Flags().GetString("output")
			if err != nil {
				log.Printf("error: %v", err)
				return err
			}

			duplicated := app.Sum().Duplicated(ctx)
			switch outputFormat {
			case "json":
				err := json.NewEncoder(os.Stdout).Encode(duplicated)
				if err != nil {
					log.Printf("error: %v", err)
					return err
				}
			case "json-i":
				encoder := json.NewEncoder(os.Stdout)
				encoder.SetIndent("", "  ")
				err := encoder.Encode(duplicated)
				if err != nil {
					log.Printf("error: %v", err)
					return err
				}
			default:
				template := "%s:\n\t%s\n\n"
				for hash, files := range duplicated {
					fmt.Printf(template, hash, strings.Join(files, "\n\t"))
				}
			}
			return nil
		},
	}

	cmd.Flags().StringP("output", "o", "text", "Output format")

	return &cmd
}
