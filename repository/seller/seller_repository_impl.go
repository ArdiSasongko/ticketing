package seller

import (
	"errors"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type SellerRepositoryImpl struct {
	db *gorm.DB
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
