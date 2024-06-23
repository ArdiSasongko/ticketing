package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/admin"
)

type BuyerService interface {
	GetBuyers(filters map[string]string, sort string, limit int, page int) ([]admin_entity.BuyerEntity, error)
	GetBuyerByID(buyerID int) (admin_entity.BuyerEntity, error)
}
