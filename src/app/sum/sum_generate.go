package sum

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/calculator"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/carlosrodriguesf/dfile/src/pkg/dbfile"
	"github.com/carlosrodriguesf/dfile/src/pkg/queue"
	"log"
	"sync"
)

func (a appImpl) Generate(ctx context.Context) error {
	var (
		mutex  sync.Mutex
		dbFile = ctx.DBFile()
		keys   = dbFile.GetFileKeys()
		calc   = calculator.New()
		q      = queue.New(10)
	)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v", r)
			dbFile.Persist()
			panic(r)
		}
	}()

	for _, file := range keys {
		file := file

		entry, _ := dbFile.GetFile(file)
		if entry.Ready {
			continue
		}

		q.Run(func() {
			hash, err := calc.Calculate(file)
			if err != nil {
				dbFile.SetFile(file, dbfile.FileEntry{
					Ready: false,
					Error: err.Error(),
				})
				return
			}

			mutex.Lock()
			dbFile.SetFile(file, dbfile.FileEntry{
				Ready: true,
				Hash:  hash,
			})
			mutex.Unlock()
		})
	}

	q.Wait()

	return dbFile.Persist()
}
