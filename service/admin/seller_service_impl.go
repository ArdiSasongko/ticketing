package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	admin_entity "github.com/ArdiSasongko/ticketing_app/model/entity/admin"
	admin_repository "github.com/ArdiSasongko/ticketing_app/repository/admin"
)

type SellerServiceImpl struct {
	sellerRepo   admin_repository.SellerRepository
	tokenUseCase helper.TokenUseCase
}

func NewSellerService(sellerRepo admin_repository.SellerRepository, tokenUseCase helper.TokenUseCase) *SellerServiceImpl {
	return &SellerServiceImpl{
		sellerRepo:   sellerRepo,
		tokenUseCase: tokenUseCase,
	}
}

func (service *SellerServiceImpl) GetSellers(filters map[string]string, sort string, limit int, page int) ([]admin_entity.SellerEntity, error) {
	result, err := service.sellerRepo.GetSellers(filters, sort, limit, page)

	if err != nil {
		return nil, err
	}

	return admin_entity.ToSellerListEntity(result), nil
}

func (service *SellerServiceImpl) GetSellerByID(sellerID int) (admin_entity.SellerEntity, error) {
	result, err := service.sellerRepo.GetSellerByID(sellerID)

	if err != nil {
		return admin_entity.SellerEntity{}, err
	}

	return admin_entity.ToSellerEntity(result), nil
}
