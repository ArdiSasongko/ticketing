package buyer_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	buyer_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	"gorm.io/gorm"
)

type OrderRepository interface {
	WithTx(builder buyer_query_builder.OrderQueryBuilder, tx *gorm.DB) OrderRepository

	ListHistory(filters map[string]string, sort string, limit int, page int) ([]domain.History, error)
	CreateHistory(history domain.History) (domain.History, error)
	UpdateHistory(history domain.History) (domain.History, error)
	CreateHistoryItem(history domain.HistoryItem) (domain.HistoryItem, error)
	GetEvent(id int) (domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	GetLatestOrder(buyerId int) (domain.History, error)
	GetHistory(historyId int) (domain.History, error)
	GetActiveHistory(buyerId int) (domain.History, error)
	DeleteHistory(history domain.History) error
	DeleteHistoryItem(historyItem domain.HistoryItem) error
}
