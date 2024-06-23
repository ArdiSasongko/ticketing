package admin_repository

import (
	"errors"
	admin_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/admin"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	adminQueryBuilder admin_query_builder.AdminQueryBuilder
	DB                *gorm.DB
}

func NewAdminRepository(adminQueryBuilder admin_query_builder.AdminQueryBuilder, db *gorm.DB) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{
		adminQueryBuilder: adminQueryBuilder,
		DB:                db,
	}
}

func (repo *AdminRepositoryImpl) GetAdmins(filters map[string]string, sort string, limit int, page int) ([]domain.Admin, error) {
	var admins []domain.Admin

	adminQueryBuilder, err := repo.adminQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := adminQueryBuilder.Find(&admins).Error
	if err1 != nil {
		return []domain.Admin{}, err1
	}
	return admins, nil
}

func (repo *AdminRepositoryImpl) GetAdminByID(adminId int) (domain.Admin, error) {
	var admin domain.Admin
	if err := repo.DB.Where("id = ?", adminId).Take(&admin).Error; err != nil {
		return domain.Admin{}, errors.New("admin not found")
	}
	return admin, nil
}
