package seller_service

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepository seller_repository.AuthRepository
	tokenUseCase   helper.TokenUseCase
}

func NewAuthService(
	authRepository seller_repository.AuthRepository,
	tokenUseCase helper.TokenUseCase,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
		tokenUseCase:   tokenUseCase,
	}
}

func (service *AuthServiceImpl) SaveSeller(request seller_web.RegisterSellerRequest) (map[string]interface{}, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}

	request.Password = string(passHash)
	newSeller := domain.Sellers{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
	}

	saveSeller, errSaveSeller := service.authRepository.SaveSeller(newSeller)
	if errSaveSeller != nil {
		return nil, errSaveSeller
	}

	data := helper.ResponseToJson{
		"id":    saveSeller.SellerID,
		"name ": saveSeller.Name,
		"email": saveSeller.Email,
	}
	return data, nil
}

func (service *AuthServiceImpl) LoginSeller(email string, password string) (map[string]interface{}, error) {
	seller, getSellerErr := service.authRepository.FindUserByEmail(email)
	if getSellerErr != nil {
		return nil, errors.New("wrong email or password")
	}

	if checkPasswordErr := bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(password)); checkPasswordErr != nil {
		return nil, errors.New("wrong email or password")
	}

	loginResponse, loginErr := helper.Login(enum.UserTypeSeller, seller.SellerID, seller.Email)
	if loginErr != nil {
		return nil, loginErr
	}

	return helper.CustomResponse{
		"token":      loginResponse["token"],
		"expires_at": loginResponse["expires_at"],
	}, nil
}

func (service *AuthServiceImpl) GetSeller(sellerId int) (seller_entity.SellerEntity, error) {
	getSeller, errGetSeller := service.authRepository.GetSeller(sellerId)

	if errGetSeller != nil {
		return seller_entity.SellerEntity{}, errGetSeller
	}

	return seller_entity.ToSellerEntity(getSeller), nil
}

func (service *AuthServiceImpl) UpdateSeller(request seller_web.UpdateSellerRequest, pathId int) (map[string]interface{}, error) {
	getSellerById, err := service.authRepository.GetSeller(pathId)
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
	}

	sellerRequest, errUpdate := service.authRepository.UpdateSeller(sellerRequest)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return helper.CustomResponse{
		"id":    sellerRequest.SellerID,
		"name":  sellerRequest.Name,
		"email": sellerRequest.Email,
	}, nil
}
