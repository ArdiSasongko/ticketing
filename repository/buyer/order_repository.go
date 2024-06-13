package buyer_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	WithTx(tx *gorm.DB) OrderRepository

	ListHistory() ([]domain.History, error)
	CreateHistory(history domain.History) (domain.History, error)
	UpdateHistory(history domain.History) (domain.History, error)
	CreateHistoryItem(history domain.HistoryItem) (domain.HistoryItem, error)
	GetEvent(id int) (domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	GetLatestOrder() (domain.History, error)
	GetHistory(historyId int) (domain.History, error)
	GetActiveHistory(buyerId int) (domain.History, error)
	DeleteHistory(history domain.History) error
	DeleteHistoryItem(historyItem domain.HistoryItem) error
}
