package admin_service

import (
	"errors"
	"strconv"
	"time"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	seller_entity "github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	admin_web "github.com/ArdiSasongko/ticketing_app/model/web/admin"
	admin_repository "github.com/ArdiSasongko/ticketing_app/repository/admin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	repo  admin_repository.AdminRepositoryInterface
	Token helper.TokenUseCase
}

func NewAdminService(repo admin_repository.AdminRepositoryInterface, token helper.TokenUseCase) *AdminService {
	return &AdminService{
		repo:  repo,
		Token: token,
	}
}

func (service *AdminService) Register(req admin_web.AdminRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newAdmin := domain.Admins{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	result, err := service.repo.Register(newAdmin)

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

func (service *AdminService) Login(email, password string) (helper.CustomResponse, error) {
	user, errUser := service.repo.GetEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	if errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errPass != nil {
		return nil, errors.New("password not match")
	}

	expiredTime := time.Now().Add(5 * time.Minute)

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(user.AdminID),
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

func (service *AdminService) GetBuyers() ([]buyer_entity.BuyerEntity, error) {
	result, err := service.repo.GetBuyers()

	if err != nil {
		return nil, err
	}

	return buyer_entity.ToBuyerEntities(result), nil
}

func (service *AdminService) GetSellers() ([]seller_entity.SellerEntity, error) {
	result, err := service.repo.GetSellers()

	if err != nil {
		return nil, err
	}

	return seller_entity.ToSellerrEntities(result), nil
}

func (service *AdminService) GetBuyerByID(buyerID int) (buyer_entity.BuyerEntity, error) {
	result, err := service.repo.GetBuyerByID(buyerID)

	if err != nil {
		return buyer_entity.BuyerEntity{}, err
	}

	return buyer_entity.ToBuyerEntity(result), nil
}

func (service *AdminService) GetSellerByID(sellerID int) (seller_entity.SellerEntity, error) {
	result, err := service.repo.GetSellerByID(sellerID)

	if err != nil {
		return seller_entity.SellerEntity{}, err
	}

	return seller_entity.ToSellerEntity(result), nil
}
