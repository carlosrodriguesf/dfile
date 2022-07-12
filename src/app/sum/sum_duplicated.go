package sum

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
)

func (a appImpl) Duplicated(ctx context.Context) map[string][]string {
	var (
		dbFile = ctx.DBFile()
		keys   = dbFile.GetFileKeys()
		sumMap = make(map[string][]string)
	)

	for _, file := range keys {
		entry := dbFile.GetFile(file)
		if sumMap[entry.Hash] == nil {
			sumMap[entry.Hash] = []string{file}
			continue
		}
		sumMap[entry.Hash] = append(sumMap[entry.Hash], file)
	}

	for key, files := range sumMap {
		if len(files) == 1 {
			delete(sumMap, key)
		}
	}

	return sumMap
}
