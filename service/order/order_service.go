package order_service

import (
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	buyer_web "github.com/ArdiSasongko/ticketing_app/model/web/buyer"
)

type OrderService interface {
	Order(id int, req buyer_web.Order) (map[string]interface{}, error)
	GenerateHistoryItem() error
	GenerateOrderNumber() (string, error)
	GetLatestOrder() ([]buyer_entity.HistoryItemEntity, error)
	PayOrder(BuyyerIDFK uint) (*buyer_entity.HistoryEntity, error)
}
