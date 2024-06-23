package admin_repository

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repo *AuthRepositoryImpl) Register(admin domain.Admin) (domain.Admin, error) {
	if err := repo.DB.Create(&admin).Error; err != nil {
		return domain.Admin{}, err
	}

	return admin, nil
}

func (repo *AuthRepositoryImpl) GetEmail(email string) (domain.Admin, error) {
	var admin domain.Admin
	if err := repo.DB.Where("email = ?", email).Take(&admin).Error; err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}

func (repo *AuthRepositoryImpl) Update(userID int, admin domain.Admin) (domain.Admin, error) {
	if err := repo.DB.Model(&domain.Admin{}).Where("id = ?", userID).Updates(admin).Error; err != nil {
		return domain.Admin{}, errors.New("failed to update admin")
	}

	return admin, nil
}

func (repo *AuthRepositoryImpl) GetByID(userID int) (domain.Admin, error) {
	var admin domain.Admin
	if err := repo.DB.Where("id = ?", userID).Take(&admin).Error; err != nil {
		return domain.Admin{}, errors.New("admin not found")
	}
	return admin, nil
}
