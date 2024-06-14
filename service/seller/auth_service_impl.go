package seller_service

import (
	"errors"
	"strconv"
	"time"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	seller_entity "github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	seller_web "github.com/ArdiSasongko/ticketing_app/model/web/seller"
	seller_repository "github.com/ArdiSasongko/ticketing_app/repository/seller"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SellerServiceImpl struct {
	repository   seller_repository.SellerRepository
	tokenUseCase helper.TokenUseCase
}

func NewSellerService(repository seller_repository.SellerRepository, token helper.TokenUseCase) *SellerServiceImpl {
	return &SellerServiceImpl{
		repository:   repository,
		tokenUseCase: token,
	}
}

func (service *SellerServiceImpl) SaveSeller(request seller_web.SellerServiceRequest) (map[string]interface{}, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}

	request.Password = string(passHash)
	newseller := domain.Sellers{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
	}

	saveSeller, errSaveSeller := service.repository.SaveSeller(newseller)
	if errSaveSeller != nil {
		return nil, errSaveSeller
	}

	data := helper.ResponseToJson{
		"name ": saveSeller.Name,
		"email": saveSeller.Email,
	}
	return data, nil
}

func (service *SellerServiceImpl) LoginSeller(email string, password string) (map[string]interface{}, error) {
	seller, err := service.repository.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(password))
	if errPass != nil {
		return nil, errors.New("password Salah")
	}

	expiredTime := time.Now().Local().Add(1 * time.Hour)

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(seller.SellerID),
		Name:  seller.Name,
		Email: seller.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest-gorm",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token, errToken := service.tokenUseCase.GenerateAccessToken(claims)
	if errToken != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return map[string]interface{}{"token": token}, nil
}

func (service *SellerServiceImpl) GetSeller(sellerId int) (seller_entity.SellerEntity, error) {
	getSeller, errGetSeller := service.repository.GetSeller(sellerId)

	if errGetSeller != nil {
		return seller_entity.SellerEntity{}, errGetSeller
	}

	return seller_entity.ToSellerEntity(getSeller), nil
}

func (service *SellerServiceImpl) UpdateSeller(request seller_web.SellerUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
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
