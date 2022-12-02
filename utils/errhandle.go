package utils

import "errors"

var (
	ErrDbErr      = errors.New("db err")
	ErrInvalid    = errors.New("invalid argument")
	ErrPermission = errors.New("permission denied")
	ErrEmpty      = errors.New("no data")
	ErrExist      = errors.New("data already exists")
	ErrNotExist   = errors.New("data does not exist")
)
