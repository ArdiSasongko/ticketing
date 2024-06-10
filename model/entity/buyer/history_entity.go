package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type HistoryEntity struct {
	Id            int       `json:"id"`
	BuyyerIDFK    int       `json:"buyer_id"`
	Number        string    `json:"number"`
	PaymentStatus string    `json:"payment_status"`
	Total         float64   `json:"total"`
	PaidAt        time.Time `json:"paid_at"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func ToHistoryEntity(history domain.History) HistoryEntity {
	return HistoryEntity{
		Id:            history.Id,
		BuyyerIDFK:    history.BuyyerIDFK,
		Number:        history.Number,
		PaymentStatus: history.PaymentStatus,
		Total:         history.Total,
		PaidAt:        history.PaidAt,
		CreatedAt:     history.CreatedAt,
		UpdatedAt:     history.UpdatedAt,
	}
}
