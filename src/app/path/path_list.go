package path

import (
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
)

func (a *appImpl) List() ([]string, error) {
	allPaths, err := a.pathRep.All()
	if err != nil {
		return nil, hlog.LogError(err)
	}

	var (
		i     = 0
		paths = make([]string, len(allPaths))
	)
	for _, path := range allPaths {
		paths[i] = path.Path
		i++
	}
	return paths, nil
}
