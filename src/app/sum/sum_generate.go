package sum

import (
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"github.com/carlosrodriguesf/dfile/src/tool/queue"
)

func (a appImpl) Generate() error {
	allFiles, err := a.fileRep.All()
	if err != nil {
		return hlog.LogError(err)
	}

	q := queue.New(10)
	for _, file := range allFiles {
		fmt.Println(file)
		if file.Checksum != "" {
			continue
		}

		file := file
		q.Run(func() {
			a.GenerateFileSum(file.Path)
		})
	}

	q.Wait()

	return nil
}

func (a *appImpl) GenerateFileSum(file string) {
	var hash, err = a.calculator.Calculate(file)

	if err != nil {
		hlog.Error(err)
		a.saveError(file, err)
		return
	}

	a.saveChecksum(file, hash)
}

func (a *appImpl) saveError(file string, err error) {
	err = a.fileRep.Save(model.File{
		Path:  file,
		Error: err.Error(),
	})
	if err != nil {
		hlog.Error(err)
	}
}

func (a *appImpl) saveChecksum(file string, hash string) {
	err := a.fileRep.Save(model.File{
		Path:     file,
		Checksum: hash,
	})
	if err != nil {
		hlog.Error(err)
	}
}
