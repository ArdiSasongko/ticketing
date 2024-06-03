package seller

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type SellerRepository interface {
	SaveSeller(seller domain.Sellers) (domain.Sellers, error)
	FindUserByEmail(email string) (*domain.Sellers, error)
}
