package errors

import (
	"testing"

	"github.com/pkg/errors"
)

func TestError(t *testing.T) {
	custErr := errors.New("error user already verify")
	err := errors.Wrap(custErr, "123")
	DecodeError(err)

}
