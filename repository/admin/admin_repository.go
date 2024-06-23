package admin_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type AdminRepository interface {
	GetAdmins(filters map[string]string, sort string, limit int, page int) ([]domain.Admin, error)
	GetAdminByID(adminId int) (domain.Admin, error)
}
