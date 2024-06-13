package buyer_entity

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"time"
)

type HistoryLiteEntity struct {
	Id            int       `json:"id"`
	Number        string    `json:"number"`
	PaymentStatus string    `json:"payment_status"`
	Total         float64   `json:"total"`
	PaidAt        time.Time `json:"paid_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToHistoryLiteEntity(history domain.History) HistoryLiteEntity {
	return HistoryLiteEntity{
		Id:            history.Id,
		Number:        history.Number,
		PaymentStatus: history.PaymentStatus,
		Total:         history.Total,
		PaidAt:        history.PaidAt,
		CreatedAt:     history.CreatedAt,
		UpdatedAt:     history.UpdatedAt,
	}
}

func ToHistoryLiteEntityCollection(histories []domain.History) []HistoryLiteEntity {
	var historyList []HistoryLiteEntity

	for _, history := range histories {
		historyList = append(historyList, ToHistoryLiteEntity(history))
	}

	return historyList
}
