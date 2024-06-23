package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/admin"
)

type AdminService interface {
	GetAdmins(filters map[string]string, sort string, limit int, page int) ([]admin_entity.AdminEntity, error)
	GetAdminByID(adminId int) (admin_entity.AdminEntity, error)
}
