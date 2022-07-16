package path

import (
	"github.com/carlosrodriguesf/dfile/src/model"
	"github.com/carlosrodriguesf/dfile/src/tool/dbm"
)

type path struct {
	model.Path
	IgnoreFolders    dbm.StringArray `json:"ignore_folders" db:"ignore_folders"`
	AcceptExtensions dbm.StringArray `json:"accept_extensions" db:"accept_extensions"`
}
