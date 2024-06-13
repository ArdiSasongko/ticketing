package buyer_entity

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type BuyerEntity struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToBuyerEntity(buyer domain.Buyers) BuyerEntity {
	return BuyerEntity{
		ID:    buyer.BuyerID,
		Name:  buyer.Name,
		Email: buyer.Email,
	}
}

func ToBuyerEntities(buyers []domain.Buyers) []BuyerEntity {
	var result []BuyerEntity
	for _, buyer := range buyers {
		result = append(result, ToBuyerEntity(buyer))
	}
	return result
}
