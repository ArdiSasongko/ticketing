package buyer

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

func (bR *BuyerRepo) Register(buyer domain.Buyers) (domain.Buyers, error) {
	if err := bR.DB.Create(&buyer).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.Buyers{}, errors.New("email already registered")
		}
	}
	return buyer, nil
}

func (bR *BuyerRepo) GetEmail(email string) (domain.Buyers, error) {
	var buyer domain.Buyers
	if err := bR.DB.Where("email = ?", email).Take(&buyer).Error; err != nil {
		return domain.Buyers{}, errors.New("email not found")
	}
	return buyer, nil
}
