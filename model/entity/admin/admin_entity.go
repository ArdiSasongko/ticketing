package admin_entity

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type AdminEntity struct {
	AdminID int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

func ToAdminEntity(admin domain.Admin) AdminEntity {
	return AdminEntity{
		AdminID: admin.ID,
		Name:    admin.Name,
		Email:   admin.Email,
	}
}

func ToAdminListEntity(admins []domain.Admin) []AdminEntity {
	var adminData []AdminEntity

	for _, admin := range admins {
		adminData = append(adminData, ToAdminEntity(admin))

	}
	return adminData
}
