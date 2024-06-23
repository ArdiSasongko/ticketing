package admin_repository

import (
	"errors"
	admin_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/admin"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type SellerRepositoryImpl struct {
	sellerQueryBuilder admin_query_builder.SellerQueryBuilder
	DB                 *gorm.DB
}

func NewSellerRepository(sellerQueryBuilder admin_query_builder.SellerQueryBuilder, db *gorm.DB) *SellerRepositoryImpl {
	return &SellerRepositoryImpl{
		sellerQueryBuilder: sellerQueryBuilder,
		DB:                 db,
	}
}

func (repo *SellerRepositoryImpl) GetSellers(filters map[string]string, sort string, limit int, page int) ([]domain.Sellers, error) {
	var sellers []domain.Sellers

	eventQueryBuilder, err := repo.sellerQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := eventQueryBuilder.Find(&sellers).Error
	if err1 != nil {
		return []domain.Sellers{}, err1
	}
	return sellers, nil
}

func (repo *SellerRepositoryImpl) GetSellerByID(sellerID int) (domain.Sellers, error) {
	var seller domain.Sellers
	if err := repo.DB.Where("id = ?", sellerID).Take(&seller).Error; err != nil {
		return domain.Sellers{}, errors.New("buyer not found")
	}
	return seller, nil
}
