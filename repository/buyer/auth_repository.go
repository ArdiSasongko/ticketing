package buyer_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type AuthRepository interface {
	GetEmail(email string) (domain.Buyer, error)
	Register(buyer domain.Buyer) (domain.Buyer, error)
	Update(userID int, buyer domain.Buyer) (domain.Buyer, error)
	GetByID(userID int) (domain.Buyer, error)
}
