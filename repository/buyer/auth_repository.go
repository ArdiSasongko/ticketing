package buyer_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type BuyerRepositoryInterface interface {
	GetEmail(email string) (domain.Buyers, error)
	Register(buyer domain.Buyers) (domain.Buyers, error)
	Update(userID int, buyer domain.Buyers) (domain.Buyers, error)
	GetByID(userID int) (domain.Buyers, error)
}
