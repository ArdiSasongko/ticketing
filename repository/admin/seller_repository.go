package admin_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type SellerRepository interface {
	GetSellers(filters map[string]string, sort string, limit int, page int) ([]domain.Sellers, error)
	GetSellerByID(sellerID int) (domain.Sellers, error)
}
