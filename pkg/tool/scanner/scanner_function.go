package scanner

import (
	"os"
	"strings"
)

func getEntryExtension(entry os.DirEntry) string {
	parts := strings.Split(entry.Name(), ".")
	return parts[len(parts)-1]
}

func acceptAllOptions(entry os.DirEntry, opts []Option) bool {
	for _, opt := range opts {
		if !opt(entry) {
			return false
		}
	}
	return true
}
