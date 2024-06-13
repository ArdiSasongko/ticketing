package buyer_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type SellerRepositoryImpl struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepositoryImpl {
	return &SellerRepositoryImpl{db: db}
}

func (repo *SellerRepositoryImpl) AddSellerBalance(seller domain.Sellers, amount float64) (domain.Sellers, error) {
	updateSellerBalanceErr := repo.db.Exec("UPDATE seller SET balance = balance + ?", amount).Error
	if updateSellerBalanceErr != nil {
		return domain.Sellers{}, updateSellerBalanceErr
	}

	getSellerErr := repo.db.Where("id = ?", seller.SellerID).First(&seller).Error
	if getSellerErr != nil {
		return domain.Sellers{}, getSellerErr
	}

	return seller, nil
}
