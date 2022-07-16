package apperrors

import "errors"

var (
	ErrPathAlreadyExists = errors.New("path already exists")
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrFileNotExists     = errors.New("file not exists")
)
