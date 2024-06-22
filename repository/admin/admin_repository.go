package admin_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type AdminRepositoryInterface interface {
	Register(admin domain.Admins) (domain.Admins, error)
	GetEmail(email string) (domain.Admins, error)
	GetBuyers() ([]domain.Buyers, error)
	GetSellers() ([]domain.Sellers, error)
	GetBuyerByID(buyerID int) (domain.Buyers, error)
	GetSellerByID(sellerID int) (domain.Sellers, error)
}
