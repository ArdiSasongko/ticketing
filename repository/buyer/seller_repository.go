package buyer_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type SellerRepository interface {
	AddSellerBalance(seller domain.Sellers, amount float64) (domain.Sellers, error)
}
