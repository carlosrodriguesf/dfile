package dbm

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
)

type StringArray []string

func (arr *StringArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal([]byte(value.(string)), arr)
	if err != nil {
		return hlog.LogError(err)
	}
	return nil
}

func (arr StringArray) Value() (driver.Value, error) {
	dt, err := json.Marshal(arr)
	if err != nil {
		return nil, hlog.LogError(err)
	}
	return string(dt), nil
}

func JSONArray(v interface{}) interface {
	driver.Valuer
	sql.Scanner
} {
	switch v := v.(type) {
	case []string:
		return (*StringArray)(&v)
	case *[]string:
		return (*StringArray)(v)
	}
	return nil
}
