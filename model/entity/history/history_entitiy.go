package history_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type HistoryEntity struct {
	ID            int        `json:"id"`
	BuyerIDFK     int        `json:"buyer_id"`
	Number        string     `json:"number"`
	PaymentStatus string     `json:"payment_status"`
	Total         float64    `json:"total"`
	PaidAt        *time.Time `json:"paid_at"`
}

func ToHistoryEntity(history domain.History) HistoryEntity {
	return HistoryEntity{
		ID:            history.Id,
		BuyerIDFK:     history.BuyerIDFK,
		Number:        history.Number,
		PaymentStatus: history.PaymentStatus,
		Total:         history.Total,
		PaidAt:        history.PaidAt,
	}
}

func ToHistoryEntities(histories []domain.History) []HistoryEntity {
	var result []HistoryEntity
	for _, history := range histories {
		result = append(result, ToHistoryEntity(history))
	}
	return result
}
