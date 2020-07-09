package common

import "errors"

var ErrNotFound = errors.New("requested item not found")


func ErrorHandler(e error) {
}
