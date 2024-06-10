package buyer_service

import (
	"errors"
	"strconv"
	"time"

	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/history"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/entity/history"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type BuyerService struct {
	Repo        buyer_repository.BuyerRepositoryInterface
	HistoryRepo history_repository.HistoryRepo
	Token       helper.TokenUseCase
}

func NewBuyerService(repo buyer_repository.BuyerRepositoryInterface, token helper.TokenUseCase, history history_repository.HistoryRepo) *BuyerService {
	return &BuyerService{
		Repo:        repo,
		HistoryRepo: history,
		Token:       token,
	}
}

func (service *BuyerService) Register(req buyer_web.BuyerRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newBuyer := domain.Buyers{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	result, err := service.Repo.Register(newBuyer)

	if err != nil {
		return nil, err
	}

	data := helper.CustomResponse{
		"name":     result.Name,
		"email":    result.Email,
		"password": result.Password,
	}

	return data, nil
}

func (service *BuyerService) Login(email, password string) (helper.CustomResponse, error) {
	user, errUser := service.Repo.GetEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	if errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errPass != nil {
		return nil, errors.New("password not match")
	}

	expiredTime := time.Now().Add(5 * time.Minute)

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(user.BuyerID),
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Echo",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := service.Token.GenerateAccessToken(claims)

	if errToken != nil {
		return nil, errToken
	}

	data := helper.CustomResponse{
		"token":      token,
		"expired_at": expiredTime.Format(time.RFC3339),
	}

	return data, nil
}

func (service *BuyerService) Update(userId int, req buyer_web.BuyerUpdateRequest) (helper.CustomResponse, error) {
	buyer, errBuyer := service.Repo.GetByID(userId)

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

	result, errUpdate := service.Repo.Update(userId, buyer)

	if errUpdate != nil {
		return nil, errUpdate
	}

	data := helper.CustomResponse{
		"name":     result.Name,
		"email":    result.Email,
		"password": result.Password,
	}

	return data, nil
}

func (service *BuyerService) GetAll() ([]buyer_entity.BuyerEntity, error) {
	result, err := service.Repo.GetList()

	if err != nil {
		return nil, err
	}

	return buyer_entity.ToBuyerEntities(result), nil
}

func (service *BuyerService) GetHistory(userId int) ([]history_entity.HistoryEntity, error) {
	result, err := service.HistoryRepo.GetHistory(userId)

	if err != nil {
		return nil, err
	}

	return history_entity.ToHistoryEntities(result), nil
}
