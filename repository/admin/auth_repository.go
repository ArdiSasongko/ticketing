package admin_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type AuthRepository interface {
	Register(admin domain.Admin) (domain.Admin, error)
	GetEmail(email string) (domain.Admin, error)
	Update(userID int, admin domain.Admin) (domain.Admin, error)
	GetByID(userID int) (domain.Admin, error)
}
