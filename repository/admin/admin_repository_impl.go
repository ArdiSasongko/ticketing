package admin_repository

import (
	"errors"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (repo *AdminRepo) Register(admin domain.Admins) (domain.Admins, error) {
	if err := repo.DB.Create(&admin).Error; err != nil {
		return domain.Admins{}, err
	}

	return admin, nil
}

func (repo *AdminRepo) GetEmail(email string) (domain.Admins, error) {
	var admin domain.Admins
	if err := repo.DB.Where("email = ?", email).Take(&admin).Error; err != nil {
		return domain.Admins{}, err
	}
	return admin, nil
}

func (repo *AdminRepo) GetBuyers() ([]domain.Buyers, error) {
	var buyers []domain.Buyers
	if err := repo.DB.Find(&buyers).Error; err != nil {
		return []domain.Buyers{}, err
	}

	return buyers, nil
}

func (repo *AdminRepo) GetSellers() ([]domain.Sellers, error) {
	var sellers []domain.Sellers
	if err := repo.DB.Find(&sellers).Error; err != nil {
		return []domain.Sellers{}, err
	}

	return sellers, nil
}

func (repo *AdminRepo) GetBuyerByID(buyerID int) (domain.Buyers, error) {
	var buyer domain.Buyers
	if err := repo.DB.Where("id = ?", buyerID).Take(&buyer).Error; err != nil {
		return domain.Buyers{}, errors.New("buyer not found")
	}
	return buyer, nil
}

func (repo *AdminRepo) GetSellerByID(sellerID int) (domain.Sellers, error) {
	var seller domain.Sellers
	if err := repo.DB.Where("id = ?", sellerID).Take(&seller).Error; err != nil {
		return domain.Sellers{}, errors.New("buyer not found")
	}
	return seller, nil
}
