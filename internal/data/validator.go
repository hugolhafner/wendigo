package data

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

func ValidateStruct(data interface{}) bool {
	if err := validate.Struct(data); err != nil {
		logrus.Warnf("failed: %v", err)
		return false
	}

	return true
}
