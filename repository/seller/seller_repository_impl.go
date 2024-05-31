package seller

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type SellerRepositoryImpl struct {
	db *gorm.DB
}

func (repo *SellerRepositoryImpl) UpdateUser(seller domain.Sellers) (domain.Sellers, error) {
	err := repo.db.Model(domain.Sellers{}).Where("user_id=?", seller.SellerID).Updates(seller).Error

	if err != nil {
		return seller, err
	}

	return seller, nil
}
