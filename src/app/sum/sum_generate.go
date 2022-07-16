package sum

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/carlosrodriguesf/dfile/src/tool/dbfile"
	"github.com/carlosrodriguesf/dfile/src/tool/queue"
	"log"
	"sync"
)

func (a appImpl) Generate(ctx context.Context) error {
	a.mutex = new(sync.Mutex)

	var (
		dbFile = ctx.DBFile()
		keys   = dbFile.GetFileKeys()
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

		a.mutex.Lock()
		entry := dbFile.GetFile(file)
		a.mutex.Unlock()

		if entry.Ready || entry.Error != "" {
			continue
		}

		q.Run(func() {
			a.GenerateFileSum(ctx, file, false)
		})
	}

	q.Wait()

	return dbFile.Persist()
}

func (a *appImpl) GenerateFileSum(ctx context.Context, file string, persist bool) {
	var (
		dbFile    = ctx.DBFile()
		hash, err = a.calculator.Calculate(file)
	)

	if a.mutex != nil {
		a.mutex.Lock()
		defer a.mutex.Unlock()
	}

	if err != nil {
		dbFile.SetFile(file, dbfile.FileEntry{
			Ready: false,
			Error: err.Error(),
		})
	}
	dbFile.SetFile(file, dbfile.FileEntry{
		Ready: true,
		Hash:  hash,
	})

	if persist {
		err := dbFile.Persist()
		if err != nil {
			log.Fatal(err)
		}
	}
}
