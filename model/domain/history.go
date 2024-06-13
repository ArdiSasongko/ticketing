package domain

import (
	"time"
)

type History struct {
	Id            int       `gorm:"column:id"`
	BuyyerIDFK    int       `gorm:"column:buyer_id"`
	Number        string    `gorm:"column:number"`
	PaymentStatus string    `gorm:"column:payment_status"`
	Total         float64   `gorm:"column:total"`
	PaidAt        time.Time `gorm:"column:paid_at"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	HistoryItems  []HistoryItem
	Buyer         Buyers `gorm:"foreignKey:buyer_id"`
}

func (History) TableName() string {
	return "history"
}
