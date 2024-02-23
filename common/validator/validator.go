package validator

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	ErrMessages() map[string]string
}

func GetErrorMsg(request Validator, err error) string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range validationErrors {
			if message, exist := request.ErrMessages()[v.Field()+"."+v.Tag()]; exist {
				return message
			}
			return v.Error()
		}
	}

	unmarshalTypeError, ok := err.(*json.UnmarshalTypeError)
	if ok {
		return fmt.Sprintf("expected field[%s] type[%s] error, got %s", unmarshalTypeError.Field, unmarshalTypeError.Type, unmarshalTypeError.Value)
	}

	return err.Error()
}
