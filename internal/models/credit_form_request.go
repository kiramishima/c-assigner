package models

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type CreditPostFormRequest struct {
	Investment int32 `json:"investment,omitempty" validate:"required,number,gt=0"`
}

// Validate the input
func (u *CreditPostFormRequest) Validate(v *validator.Validate) error {
	err := v.Struct(u)
	if err != nil {

		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			var out error
			err = err.(validator.ValidationErrors)
			for _, fe := range ve {
				out = errors.New(msgForTag(fe.Tag()))
			}
			return out
		}
	}
	return nil
}

// only for this case
func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "gt":
		return "The value needs to be more than 0 and non negative"
	}
	return ""
}
