package admin_entity

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type SellerEntity struct {
	SellerID int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToSellerEntity(seller domain.Sellers) SellerEntity {
	return SellerEntity{
		SellerID: seller.SellerID,
		Name:     seller.Name,
		Email:    seller.Email,
	}
}

func ToSellerListEntity(sellers []domain.Sellers) []SellerEntity {
	var result []SellerEntity
	for _, seller := range sellers {
		result = append(result, ToSellerEntity(seller))
	}
	return result
}
