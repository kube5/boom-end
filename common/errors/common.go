package errors

import "github.com/pkg/errors"

var (
	ErrRequestParameter         = errors.New("error request parameter")
	ErrRequestParameterAllEmpty = errors.New("error request parameter all empty")

	ErrRedisNotFound = errors.New("redis not found")

	ErrOpenAPIAuth = errors.New("error open admin collection")

	ErrESNotFound = errors.New("ES not found")

	ErrWrongConfig = errors.New("wrong config")

	ErrDuplicateSubscription = errors.New("error duplicate subscription")

	ErrInternal = errors.New("error internal")
)
