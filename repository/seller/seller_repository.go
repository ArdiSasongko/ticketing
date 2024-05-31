package seller

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type SellerRepository interface {
	UpdateSeller(seller domain.Sellers) (domain.Sellers, error)
	GetSeller(Id int) (domain.Sellers, error)
}
