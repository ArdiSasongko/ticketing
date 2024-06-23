package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	admin_entity "github.com/ArdiSasongko/ticketing_app/model/entity/admin"
	admin_repository "github.com/ArdiSasongko/ticketing_app/repository/admin"
)

type BuyerServiceImpl struct {
	buyerRepo    admin_repository.BuyerRepository
	tokenUseCase helper.TokenUseCase
}

func NewBuyerService(buyerRepo admin_repository.BuyerRepository, tokenUseCase helper.TokenUseCase) *BuyerServiceImpl {
	return &BuyerServiceImpl{
		buyerRepo:    buyerRepo,
		tokenUseCase: tokenUseCase,
	}
}

func (service *BuyerServiceImpl) GetBuyers(filters map[string]string, sort string, limit int, page int) ([]admin_entity.BuyerEntity, error) {
	result, err := service.buyerRepo.GetBuyers(filters, sort, limit, page)

	if err != nil {
		return nil, err
	}

	return admin_entity.ToBuyerListEntity(result), nil
}

func (service *BuyerServiceImpl) GetBuyerByID(buyerID int) (admin_entity.BuyerEntity, error) {
	result, err := service.buyerRepo.GetBuyerByID(buyerID)

	if err != nil {
		return admin_entity.BuyerEntity{}, err
	}

	return admin_entity.ToBuyerEntity(result), nil
}
