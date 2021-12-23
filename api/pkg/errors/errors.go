package errors

import (
	"github.com/pkg/errors"
)

const (
	KindNotExists = iota
	KindAlreadyExists
)

var _ error = (*kindError)(nil)

type kindError struct {
	err  error
	kind int
}

func NewKindedError(kind int, msgf string, args ...interface{}) error {
	return &kindError{
		kind: kind,
		err:  errors.Errorf(msgf, args...),
	}
}

func (e *kindError) Error() string {
	return e.err.Error()
}
func (e *kindError) Kind() int {
	return e.kind
}

func IsNotExistsError(err error) bool {
	return GetErrorKind(err) == KindNotExists
}

func GetErrorKind(err error) int {
	switch e := err.(type) {
	case *kindError:
		return e.kind
	}
	return -1
}
