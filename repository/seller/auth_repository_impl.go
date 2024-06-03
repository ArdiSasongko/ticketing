package seller

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
