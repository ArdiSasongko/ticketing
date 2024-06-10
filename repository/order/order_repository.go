package order_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
)

type OrderRepositoryInterface interface {
	Order(id int, qty domain.Event) (domain.Event, error)
	GetByID(ID int) (domain.Event, error)
	CopyEvent() ([]buyer_entity.EventEntity, error)
	PasteHistory(destinations []buyer_entity.HistoryItemEntity) error
	GetLatestOrder() (*buyer_entity.HistoryEntity, error)
	GetOrderByID(BuyyerIDFK uint) (*buyer_entity.HistoryEntity, error)
	UpdateOrder(order *buyer_entity.HistoryEntity) error
}
