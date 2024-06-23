package admin_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type BuyerRepository interface {
	GetBuyers(filters map[string]string, sort string, limit int, page int) ([]domain.Buyer, error)
	GetBuyerByID(buyerID int) (domain.Buyer, error)
}
