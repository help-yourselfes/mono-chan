package customErrors

import (
	"errors"
)

var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNoPasswordSet    = errors.New("no password set")
	ErrIncorectPassword = errors.New("password is incorect")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrUnableToDelete   = errors.New("unable to delete")
	ErrUpdateToUpdate   = errors.New("unable to update")
)
