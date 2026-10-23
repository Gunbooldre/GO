package storage

import (
	"errors"
)

var (
	ErrURLNotFound = errors.New("Url not found")
	ErrUrlExist    = errors.New("Url already exists")
)
