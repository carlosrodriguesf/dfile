package sum

import "github.com/carlosrodriguesf/dfile/src/tool/hlog"

func (a appImpl) Duplicated() (map[string][]string, error) {
	allFiles, err := a.fileRep.All()
	if err != nil {
		return nil, hlog.LogError(err)
	}

	sumMap := make(map[string][]string)
	for _, file := range allFiles {
		if sumMap[file.Checksum] == nil {
			sumMap[file.Checksum] = []string{file.Path}
			continue
		}
		sumMap[file.Checksum] = append(
			sumMap[file.Checksum],
			file.Path,
		)
	}

	for key, files := range sumMap {
		if len(files) == 1 {
			delete(sumMap, key)
		}
	}

	return sumMap, nil
}
