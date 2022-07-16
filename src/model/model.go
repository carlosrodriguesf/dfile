package model

type (
	Path struct {
		Path             string   `json:"path"`
		Enabled          bool     `json:"enabled"`
		IgnoreFolders    []string `json:"ignore_folders"`
		AcceptExtensions []string `json:"accept_extensions"`
	}

	File struct {
		Path     string `json:"path"`
		Checksum string `json:"checksum"`
		Error    string `json:"error"`
	}
)
