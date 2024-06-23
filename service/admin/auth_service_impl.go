package admin_service

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	admin_entity "github.com/ArdiSasongko/ticketing_app/model/entity/admin"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/ArdiSasongko/ticketing_app/model/web/admin"
	"github.com/ArdiSasongko/ticketing_app/repository/admin"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepo     admin_repository.AuthRepository
	tokenUseCase helper.TokenUseCase
}

func NewAuthService(authRepo admin_repository.AuthRepository, tokenUseCase helper.TokenUseCase) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepo:     authRepo,
		tokenUseCase: tokenUseCase,
	}
}

func (service *AuthServiceImpl) Register(req admin_web.RegisterAdminRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newAdmin := domain.Admin{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	result, err := service.authRepo.Register(newAdmin)

	if err != nil {
		return nil, err
	}

	return helper.CustomResponse{
		"id":    result.ID,
		"name":  result.Name,
		"email": result.Email,
	}, nil
}

func (service *AuthServiceImpl) Login(req admin_web.LoginAdminRequest) (helper.CustomResponse, error) {
	admin, getAdminErr := service.authRepo.GetEmail(req.Email)
	if getAdminErr != nil {
		return nil, errors.New("wrong email or password")
	}

	if checkPasswordErr := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); checkPasswordErr != nil {
		return nil, errors.New("wrong email or password")
	}

	loginResponse, loginErr := helper.Login(enum.UserTypeAdmin, admin.ID, admin.Email)
	if loginErr != nil {
		return nil, loginErr
	}

	return helper.CustomResponse{
		"token":      loginResponse["token"],
		"expires_at": loginResponse["expires_at"],
	}, nil
}

func (service *AuthServiceImpl) Update(userId int, req admin_web.UpdateAdminRequest) (helper.CustomResponse, error) {
	admin, errBuyer := service.authRepo.GetByID(userId)

	if errBuyer != nil {
		return nil, errBuyer
	}

	if req.Name != "" {
		admin.Name = req.Name
	}

	if req.Email != "" {
		admin.Email = req.Email
	}

	if req.Password != "" {
		passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

		if errHash != nil {
			return nil, errHash
		}

		admin.Password = string(passHash)
	}

	result, errUpdate := service.authRepo.Update(userId, admin)

	if errUpdate != nil {
		return nil, errUpdate
	}

	data := helper.CustomResponse{
		"id":    result.ID,
		"name":  result.Name,
		"email": result.Email,
	}

	return data, nil
}

func (service *AuthServiceImpl) ViewMe(adminId int) (admin_entity.AdminEntity, error) {
	admin, err := service.authRepo.GetByID(adminId)
	if err != nil {
		return admin_entity.AdminEntity{}, err
	}

	return admin_entity.ToAdminEntity(admin), nil
}
