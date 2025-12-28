package customErrors

import (
	"database/sql"
	"errors"
)

var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUnableToDelete  = errors.New("unable to delete")
	ErrUpdateToUpdate  = errors.New("unable to update")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound) || err == sql.ErrNoRows
}
