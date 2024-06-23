package admin_service

import (
	admin_entity "github.com/ArdiSasongko/ticketing_app/model/entity/admin"
)

type SellerService interface {
	GetSellers(filters map[string]string, sort string, limit int, page int) ([]admin_entity.SellerEntity, error)
	GetSellerByID(sellerID int) (admin_entity.SellerEntity, error)
}
