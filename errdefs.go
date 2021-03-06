package perm

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"
)

type ErrNotFound struct {
	model string
}

func NewErrNotFound(model string) ErrNotFound {
	return ErrNotFound{
		model: model,
	}
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", err.model)
}

type ErrAlreadyExists struct {
	model string
}

func NewErrAlreadyExists(model string) ErrAlreadyExists {
	return ErrAlreadyExists{
		model: model,
	}
}

func (err ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists", err.model)
}

func NewErrorFromStatus(s *status.Status) error {
	message := fmt.Sprintf("%s", s.Message())
	return errors.New(message)
}
