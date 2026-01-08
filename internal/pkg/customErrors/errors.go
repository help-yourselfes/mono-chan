package customErrors

import (
	"errors"
)

var (
	ErrNotImplemented   = errors.New("not yet implemented")
	ErrInvalidInput     = errors.New("invalid input")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNoPasswordSet    = errors.New("no password set")
	ErrIncorectPassword = errors.New("password is incorect")
	ErrUnableToDelete   = errors.New("unable to delete")
	ErrPostIsRoot       = errors.New("the post is a thread itself")
	ErrUnableToUpdate   = errors.New("unable to update")
)
