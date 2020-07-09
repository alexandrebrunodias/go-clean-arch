package common

import "errors"

var (
	ErrNotFound = errors.New("requested item not found")
	ErrInternal = errors.New("internal server error")
)
