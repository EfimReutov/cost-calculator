package store

import "errors"

var (
	ErrNilModel = errors.New("model is nil")
	ErrZeroId   = errors.New("id is zero")
	ErrNoRow    = errors.New("no row")
)
