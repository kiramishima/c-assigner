package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditPostFormRequest_Validate(t *testing.T) {
	v := validator.New()
	t.Run("Valid", func(t *testing.T) {
		var item = CreditPostFormRequest{Investment: 300}
		assert.NoError(t, item.Validate(v))
	})

	t.Run("No Valid - required", func(t *testing.T) {
		var item = CreditPostFormRequest{}
		assert.Error(t, item.Validate(v))
		assert.EqualError(t, item.Validate(v), "This field is required")
	})

	t.Run("No Valid - greter than 0 and non negative", func(t *testing.T) {
		var item = CreditPostFormRequest{Investment: -2000}
		assert.Error(t, item.Validate(v))
		assert.EqualError(t, item.Validate(v), "The value needs to be more than 0 and non negative")
	})
}
