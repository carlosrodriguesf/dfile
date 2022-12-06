package model

type (
	Path struct {
		Path               string
		Enabled            bool
		IgnoredFolderNames []string
		AcceptExtensions   []string
	}
	File struct {
		Path     string
		Checksum string
		Error    string
	}
)
