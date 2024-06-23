package admin_repository

import (
	"errors"
	admin_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/admin"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type BuyerRepositoryImpl struct {
	buyerQueryBuilder admin_query_builder.BuyerQueryBuilder
	DB                *gorm.DB
}

func NewBuyerRepository(buyerQueryBuilder admin_query_builder.BuyerQueryBuilder, db *gorm.DB) *BuyerRepositoryImpl {
	return &BuyerRepositoryImpl{
		buyerQueryBuilder: buyerQueryBuilder,
		DB:                db,
	}
}

func (repo *BuyerRepositoryImpl) GetBuyers(filters map[string]string, sort string, limit int, page int) ([]domain.Buyer, error) {
	var buyers []domain.Buyer

	buyerQueryBuilder, err := repo.buyerQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := buyerQueryBuilder.Find(&buyers).Error
	if err1 != nil {
		return []domain.Buyer{}, err1
	}
	return buyers, nil
}

func (repo *BuyerRepositoryImpl) GetBuyerByID(buyerID int) (domain.Buyer, error) {
	var buyer domain.Buyer
	if err := repo.DB.Where("id = ?", buyerID).Take(&buyer).Error; err != nil {
		return domain.Buyer{}, errors.New("buyer not found")
	}
	return buyer, nil
}
