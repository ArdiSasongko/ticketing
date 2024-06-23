package buyer_repository

import (
	"errors"
	"strings"

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

func (repo *AuthRepositoryImpl) Register(buyer domain.Buyer) (domain.Buyer, error) {
	if err := repo.DB.Create(&buyer).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.Buyer{}, errors.New("email already registered")
		}
	}
	return buyer, nil
}

func (repo *AuthRepositoryImpl) GetEmail(email string) (domain.Buyer, error) {
	var buyer domain.Buyer
	if err := repo.DB.Where("email = ?", email).Take(&buyer).Error; err != nil {
		return domain.Buyer{}, errors.New("email not found")
	}
	return buyer, nil
}

func (repo *AuthRepositoryImpl) Update(userID int, buyer domain.Buyer) (domain.Buyer, error) {
	if err := repo.DB.Model(&domain.Buyer{}).Where("id = ?", userID).Updates(buyer).Error; err != nil {
		return domain.Buyer{}, errors.New("failed to update buyer")
	}

	return buyer, nil
}

func (repo *AuthRepositoryImpl) GetByID(userID int) (domain.Buyer, error) {
	var buyer domain.Buyer
	if err := repo.DB.Where("id = ?", userID).Take(&buyer).Error; err != nil {
		return domain.Buyer{}, errors.New("buyer not found")
	}
	return buyer, nil
}
