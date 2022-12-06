package scanner

import (
	"github.com/carlosrodriguesf/dfile/pkg/tool/array"
	"os"
)

func WithValidExtensions(extensions ...string) Option {
	validExtensions := array.ToBoolMap(extensions)
	return func(entry os.DirEntry) bool {
		if entry.IsDir() {
			return true
		}
		if validExtensions[getEntryExtension(entry)] {
			return true
		}
		return false
	}
}

func WithIgnoredNames(names ...string) Option {
	invalidNames := array.ToBoolMap(names)
	return func(entry os.DirEntry) bool {
		if invalidNames[entry.Name()] {
			return false
		}
		return true
	}
}
