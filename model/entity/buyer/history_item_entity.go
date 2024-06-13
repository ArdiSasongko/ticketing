package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type HistoryItemEntity struct {
	Id        int         `json:"id"`
	Price     float64     `json:"price"`
	Qty       int         `json:"qty"`
	Subtotal  float64     `json:"subtotal"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Event     EventEntity `json:"event"`
}

func ToHistoryItemEntity(history domain.HistoryItem) HistoryItemEntity {
	return HistoryItemEntity{
		Id:        history.Id,
		Price:     history.Price,
		Qty:       history.Qty,
		Subtotal:  history.Subtotal,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
		Event:     ToEventEntity(history.Event),
	}
}

func ToHistoryItemEntityCollection(historyItems []domain.HistoryItem) []HistoryItemEntity {
	var historyItemList []HistoryItemEntity

	for _, historyItem := range historyItems {
		historyItemList = append(historyItemList, ToHistoryItemEntity(historyItem))
	}

	return historyItemList
}
