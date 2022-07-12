package dbfile

type (
	PathEntry struct {
		AcceptExtensions []string `json:"accepet_extensions"`
		IgnoreFolders    []string `json:"ignore_folders"`
	}
	FileEntry struct {
		Ready bool   `json:"ready"`
		Hash  string `json:"content,omitempty"`
		Error string `json:"error,omitempty"`
	}
	data struct {
		Version int                  `json:"version"`
		Paths   map[string]PathEntry `json:"paths"`
		Files   map[string]FileEntry `json:"files"`
	}
)
