package seller

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type SellerRepositoryImpl struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepositoryImpl {
	return &SellerRepositoryImpl{db: db}
}

func (repo *SellerRepositoryImpl) SaveSeller(seller domain.Sellers) (domain.Sellers, error) {
	err := repo.db.Create(&seller).Error
	if err != nil {
		return domain.Sellers{}, err
	}
	return seller, nil
}

func (repo *SellerRepositoryImpl) FindUserByEmail(email string) (*domain.Sellers, error) {
	seller := new(domain.Sellers)
	if err := repo.db.Where("email = ?", email).Take(seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (repo *SellerRepositoryImpl) GetSeller(Id int) (domain.Sellers, error) {
	var sellerData domain.Sellers

	err := repo.db.First(&sellerData, "user_id = ?", Id).Error

	if err != nil {
		return domain.Sellers{}, errors.New("user tidak ditemukan")
	}

	return sellerData, nil
}

func (repo *SellerRepositoryImpl) UpdateUser(seller domain.Sellers) (domain.Sellers, error) {
	err := repo.db.Model(domain.Sellers{}).Where("user_id=?", seller.SellerID).Updates(seller).Error

	if err != nil {
		return seller, err
	}

	return seller, nil
}
