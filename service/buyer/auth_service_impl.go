package buyer_service

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepository buyer_repository.AuthRepository
	tokenUseCase   helper.TokenUseCase
}

func NewAuthService(
	authRepository buyer_repository.AuthRepository,
	tokenUseCase helper.TokenUseCase,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
		tokenUseCase:   tokenUseCase,
	}
}

func (service *AuthServiceImpl) Register(req buyer_web.RegisterBuyerRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newBuyer := domain.Buyer{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	result, err := service.authRepository.Register(newBuyer)

	if err != nil {
		return nil, err
	}

	return helper.CustomResponse{
		"id":    result.BuyerID,
		"name":  result.Name,
		"email": result.Email,
	}, nil
}

func (service *AuthServiceImpl) Login(email, password string) (helper.CustomResponse, error) {
	buyer, getBuyerErr := service.authRepository.GetEmail(email)
	if getBuyerErr != nil {
		return nil, errors.New("wrong email or password")
	}

	if checkPasswordErr := bcrypt.CompareHashAndPassword([]byte(buyer.Password), []byte(password)); checkPasswordErr != nil {
		return nil, errors.New("wrong email or password")
	}

	loginResponse, loginErr := helper.Login(enum.UserTypeBuyer, buyer.BuyerID, buyer.Email)
	if loginErr != nil {
		return nil, loginErr
	}

	return helper.CustomResponse{
		"token":      loginResponse["token"],
		"expires_at": loginResponse["expires_at"],
	}, nil
}

func (service *AuthServiceImpl) Update(userId int, req buyer_web.UpdateBuyerRequest) (helper.CustomResponse, error) {
	buyer, errBuyer := service.authRepository.GetByID(userId)

	if errBuyer != nil {
		return nil, errBuyer
	}

	if req.Name != "" {
		buyer.Name = req.Name
	}

	if req.Email != "" {
		buyer.Email = req.Email
	}

	if req.Password != "" {
		passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

		if errHash != nil {
			return nil, errHash
		}

		buyer.Password = string(passHash)
	}

	result, errUpdate := service.authRepository.Update(userId, buyer)

	if errUpdate != nil {
		return nil, errUpdate
	}

	data := helper.CustomResponse{
		"id":    result.BuyerID,
		"name":  result.Name,
		"email": result.Email,
	}

	return data, nil
}

func (service *AuthServiceImpl) ViewMe(userId int) (buyer_entity.BuyerEntity, error) {
	buyer, err := service.authRepository.GetByID(userId)
	if err != nil {
		return buyer_entity.BuyerEntity{}, err
	}

	return buyer_entity.ToBuyerEntity(buyer), nil
}
