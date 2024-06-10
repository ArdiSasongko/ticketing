package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type HistoryItemEntity struct {
	Id          int     `gorm:"column:id"`
	HistoryIDFK int     `gorm:"column:history_id"`
	EventIDFK   int     `gorm:"column:event_id"`
	Price       float64 `gorm:"column:price"`
	Qty         int     `gorm:"column:qty"`
	Subtotal    float64 `gorm:"column:subtotal"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToHistoryItemEntity(history domain.HistoryItem) HistoryItemEntity {
	return HistoryItemEntity{
		Id:          history.Id,
		HistoryIDFK: history.HistoryIDFK,
		EventIDFK:   history.EventIDFK,
		Price:       history.Price,
		Qty:         history.Qty,
		Subtotal:    history.Subtotal,
		CreatedAt:   history.CreatedAt,
		UpdatedAt:   history.UpdatedAt,
	}
}
