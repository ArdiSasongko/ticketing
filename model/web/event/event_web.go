package web

import "gorm.io/gorm"

type EventDeleteServiceRequest struct {
	UserID  int    `json:"user_id"`
	Name    string `validate:"required" json:"name"`
	Email   string `validate:"required,email" json:"email"`
	Deleted gorm.DeletedAt
}
