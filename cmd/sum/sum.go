package sum

import (
	"github.com/carlosrodriguesf/dfile/pkg/calculator"
	"github.com/carlosrodriguesf/dfile/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/pkg/queue"
	"github.com/spf13/cobra"
	"sync"
)

func generate(dbFile dbfile.DBFile) *cobra.Command {
	cmd := cobra.Command{
		Use:   "generate",
		Short: "generate sum for scanner paths",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				mutex sync.Mutex
				keys  = dbFile.Keys()
				calc  = calculator.New()
				q     = queue.New(500)
			)

			for _, file := range keys {
				file := file

				entry, _ := dbFile.Get(file)
				if entry.Ready {
					continue
				}

				q.Run(func() {
					hash, err := calc.Calculate(file)
					if err != nil {
						dbFile.Set(file, dbfile.Entry{
							Ready: false,
							Error: err.Error(),
						})
						return
					}

					mutex.Lock()
					dbFile.Set(file, dbfile.Entry{
						Ready: true,
						Hash:  hash,
					})
					mutex.Unlock()
				})
			}

			q.Wait()

			return dbFile.Persist()
		},
	}

	return &cmd
}

func Sum(dbFile dbfile.DBFile) *cobra.Command {
	pathCmd := cobra.Command{
		Use:   "sum [generate]",
		Short: "generate sum for scanned paths",
		Args:  cobra.ExactArgs(1),
	}
	pathCmd.AddCommand(generate(dbFile))
	return &pathCmd
}
