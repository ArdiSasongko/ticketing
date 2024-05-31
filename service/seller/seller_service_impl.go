package seller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	web "github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
)

type SellerServiceImpl struct {
	repository seller.SellerRepository
}

func (service *SellerServiceImpl) GetSeller(sellerId int) (entity.SellerEntity, error) {
	getSeller, errGetSeller := service.repository.GetSeller(sellerId)

	if errGetSeller != nil {
		return entity.SellerEntity{}, errGetSeller
	}
	return entity.ToSellerEntity(getSeller.SellerID, getSeller.Name, getSeller.Email), nil
}

func (service *SellerServiceImpl) UpdateSeller(request web.SellerUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
	getSellerById, err := service.repository.GetSeller(pathId)
	if err != nil {
		return nil, err
	}

	if request.Name == "" {
		request.Name = getSellerById.Name
	}

	if request.Email == "" {
		request.Email = getSellerById.Email
	}

	sellerRequest := domain.Sellers{
		SellerID: pathId,
		Name:     request.Name,
		Email:    request.Email,
		Password: getSellerById.Password,
	}

	sellerRequest, errUpdate := service.repository.UpdateSeller(sellerRequest)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return helper.CustomResponse{"name": sellerRequest.Name, "email": sellerRequest.Email}, nil
}
