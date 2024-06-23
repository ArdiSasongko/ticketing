package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/entity/admin"
	"github.com/ArdiSasongko/ticketing_app/repository/admin"
)

type AdminServiceImpl struct {
	adminRepo    admin_repository.AdminRepository
	tokenUseCase helper.TokenUseCase
}

func NewAdminService(adminRepo admin_repository.AdminRepository, tokenUseCase helper.TokenUseCase) *AdminServiceImpl {
	return &AdminServiceImpl{
		adminRepo:    adminRepo,
		tokenUseCase: tokenUseCase,
	}
}

func (service *AdminServiceImpl) GetAdmins(filters map[string]string, sort string, limit int, page int) ([]admin_entity.AdminEntity, error) {
	result, err := service.adminRepo.GetAdmins(filters, sort, limit, page)

	if err != nil {
		return nil, err
	}

	return admin_entity.ToAdminListEntity(result), nil
}

func (service *AdminServiceImpl) GetAdminByID(adminId int) (admin_entity.AdminEntity, error) {
	result, err := service.adminRepo.GetAdminByID(adminId)

	if err != nil {
		return admin_entity.AdminEntity{}, err
	}

	return admin_entity.ToAdminEntity(result), nil
}
