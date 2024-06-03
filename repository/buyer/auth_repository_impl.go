package buyer_repository

import (
	"errors"
	"strings"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type BuyerRepo struct {
	DB *gorm.DB
}

func NewBuyerRepo(db *gorm.DB) *BuyerRepo {
	return &BuyerRepo{
		DB: db,
	}
}

func (repo *BuyerRepo) Register(buyer domain.Buyers) (domain.Buyers, error) {
	if err := repo.DB.Create(&buyer).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.Buyers{}, errors.New("email already registered")
		}
	}
	return buyer, nil
}

func (repo *BuyerRepo) GetEmail(email string) (domain.Buyers, error) {
	var buyer domain.Buyers
	if err := repo.DB.Where("email = ?", email).Take(&buyer).Error; err != nil {
		return domain.Buyers{}, errors.New("email not found")
	}
	return buyer, nil
}

func (repo *BuyerRepo) Update(userID int, buyer domain.Buyers) (domain.Buyers, error) {
	if err := repo.DB.Model(&domain.Buyers{}).Where("id = ?", userID).Updates(buyer).Error; err != nil {
		return domain.Buyers{}, errors.New("failed to update buyer")
	}

	return buyer, nil
}

func (repo *BuyerRepo) GetByID(userID int) (domain.Buyers, error) {
	var buyer domain.Buyers
	if err := repo.DB.Where("id = ?", userID).Take(&buyer).Error; err != nil {
		return domain.Buyers{}, errors.New("buyer not found")
	}
	return buyer, nil
}
