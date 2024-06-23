package seller_repository

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (repo *AuthRepositoryImpl) SaveSeller(seller domain.Sellers) (domain.Sellers, error) {
	err := repo.db.Create(&seller).Error
	if err != nil {
		return domain.Sellers{}, err
	}
	return seller, nil
}

func (repo *AuthRepositoryImpl) FindUserByEmail(email string) (*domain.Sellers, error) {
	seller := new(domain.Sellers)
	if err := repo.db.Where("email = ?", email).Take(seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (repo *AuthRepositoryImpl) GetSeller(Id int) (domain.Sellers, error) {
	var sellerData domain.Sellers

	err := repo.db.First(&sellerData, "id = ?", Id).Error

	if err != nil {
		return domain.Sellers{}, errors.New("user tidak ditemukan")
	}

	return sellerData, nil
}

func (repo *AuthRepositoryImpl) UpdateSeller(seller domain.Sellers) (domain.Sellers, error) {
	err := repo.db.Model(domain.Sellers{}).Where("id = ?", seller.SellerID).Updates(seller).Error

	if err != nil {
		return seller, err
	}

	return seller, nil
}
