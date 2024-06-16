package seller_web

import (
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/go-playground/validator/v10"
)

type CreateEventsRequest struct {
	SellerID int     `json:"seller_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Date     string  `json:"date" validate:"required"`
	Location string  `json:"location" validate:"required"`
	Qty      int     `json:"qty" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UserUpdateServiceRequest struct {
	SellerID int     `json:"seller_id"`
	Name     string  `json:"name" validate:"required"`
	Date     string  `json:"date" validate:"required"`
	Location string  `json:"location" validate:"required"`
	Qty      int     `json:"qty" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UpdateEventStatusRequest struct {
	Status string `json:"status" validate:"required,event_status_enum"`
}

func EventStatusEnum(fl validator.FieldLevel) bool {
	validRoles := []string{enum.EventStatusActive, enum.EventStatusInactive, enum.EventStatusClosed}
	role := fl.Field().String()
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}
