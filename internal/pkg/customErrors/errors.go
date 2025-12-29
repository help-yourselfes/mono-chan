package customErrors

import (
	"database/sql"
	"errors"
)

var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNoPasswordSet    = errors.New("no password set")
	ErrIncorectPassword = errors.New("password is incorect")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound) || err == sql.ErrNoRows
}
