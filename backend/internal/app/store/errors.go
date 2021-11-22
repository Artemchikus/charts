package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoTinkoffKey   = errors.New("tinkoff key not found")
)
