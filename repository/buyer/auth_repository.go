package buyer

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type BuyerRepositoryInterface interface {
	GetEmail(email string) (domain.Buyers, error)
	Register(buyer domain.Buyers) (domain.Buyers, error)
}
