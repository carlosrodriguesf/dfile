package path

import (
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"path/filepath"
)

func (a appImpl) Remove(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return hlog.LogError(err)
	}

	err = a.pathRep.Remove(path)
	if err != nil {
		return hlog.LogError(err)
	}

	return nil
}
