package helper

import (
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	validatorImpl := validator.New()
	err := validatorImpl.RegisterValidation("event_status_enum", EventStatusEnum)
	if err != nil {
		return nil
	}
	return &CustomValidator{validator: validatorImpl}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func EventStatusEnum(fl validator.FieldLevel) bool {
	validRoles := []string{string(enum.EventStatusActive), string(enum.EventStatusInactive), string(enum.EventStatusClosed)}
	role := fl.Field().String()
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}
