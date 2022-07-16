package model

type (
	Path struct {
		UID              int      `json:"uid" db:"uid"`
		Path             string   `json:"path" db:"path"`
		Enabled          bool     `json:"enabled" db:"enabled"`
		IgnoreFolders    []string `json:"ignore_folders" db:"ignore_folders"`
		AcceptExtensions []string `json:"accept_extensions" db:"accept_extensions"`
	}

	File struct {
		UID         int    `json:"uid" db:"uid"`
		Path        string `json:"path" db:"path"`
		ContentType string `json:"content_type" db:"content_type"`
		Checksum    string `json:"checksum" db:"checksum"`
		Error       string `json:"error" db:"error"`
	}
)
