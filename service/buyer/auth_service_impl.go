package buyer

import (
	"errors"
	"strconv"
	"time"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	buyerreq "github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type BuyerService struct {
	Repo  buyer.BuyerRepositoryInterface
	Token helper.TokenUseCase
}

func NewBuyerService(repo buyer.BuyerRepositoryInterface, token helper.TokenUseCase) *BuyerService {
	return &BuyerService{
		Repo:  repo,
		Token: token,
	}
}

func (bS *BuyerService) Register(req buyerreq.BuyerRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newBuyer := domain.Buyers{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	result, err := bS.Repo.Register(newBuyer)

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

func (bS *BuyerService) Login(email, password string) (helper.CustomResponse, error) {
	user, errUser := bS.Repo.GetEmail(email)

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

	token, errToken := bS.Token.GenerateAccessToken(claims)

	if errToken != nil {
		return nil, errToken
	}

	data := helper.CustomResponse{
		"token":      token,
		"expired_at": expiredTime.Format(time.RFC3339),
	}

	return data, nil
}

func (bS *BuyerService) Update(userId int, req buyerreq.BuyerUpdateRequest) (helper.CustomResponse, error) {
	buyer, errBuyer := bS.Repo.GetByID(userId)

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

	result, errUpdate := bS.Repo.Update(userId, buyer)

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
