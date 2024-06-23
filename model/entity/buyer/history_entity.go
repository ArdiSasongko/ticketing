package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type HistoryEntity struct {
	Id                  int                 `json:"id"`
	Number              string              `json:"number"`
	PaymentStatus       string              `json:"payment_status"`
	Total               float64             `json:"total"`
	PaidAt              *time.Time          `json:"paid_at"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	HistoryItemEntities []HistoryItemEntity `json:"history_items"`
}

func ToHistoryEntity(history domain.History) HistoryEntity {
	return HistoryEntity{
		Id:                  history.Id,
		Number:              history.Number,
		PaymentStatus:       history.PaymentStatus,
		Total:               history.Total,
		PaidAt:              history.PaidAt,
		CreatedAt:           history.CreatedAt,
		UpdatedAt:           history.UpdatedAt,
		HistoryItemEntities: ToHistoryItemEntityCollection(history.HistoryItems),
	}
}

func ToHistoryEntityCollection(histories []domain.History) []HistoryEntity {
	var historyList []HistoryEntity

	for _, history := range histories {
		historyList = append(historyList, ToHistoryEntity(history))
	}

	return historyList
}
