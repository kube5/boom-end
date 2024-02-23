package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrExceedMaxFilesSize      = errors.New("error exceeds max file size")
	ErrGenPresignUrlNotAllowed = errors.New("error not allowed to generate presign url")
	ErrUnknownPresignUrlType   = errors.New("error unknown presign url type")
	ErrUnsupportedFileFormat   = errors.New("error unsupported file format")
	ErrInvalidProjectId        = errors.New("error invalid project id")
)
