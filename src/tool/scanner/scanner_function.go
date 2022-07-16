package scanner

import "strings"

func getExtension(name string) string {
	splittedName := strings.Split(name, ".")
	return splittedName[len(splittedName)-1]
}

func mapStringArray(arr []string) map[string]bool {
	stringMap := make(map[string]bool)
	for _, str := range arr {
		stringMap[str] = true
	}
	return stringMap
}
