package seller_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type AuthRepository interface {
	SaveSeller(seller domain.Sellers) (domain.Sellers, error)
	FindUserByEmail(email string) (*domain.Sellers, error)
	UpdateSeller(seller domain.Sellers) (domain.Sellers, error)
	GetSeller(Id int) (domain.Sellers, error)
}
