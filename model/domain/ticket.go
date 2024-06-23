package domain

import (
	"time"
)

type Ticket struct {
	Id        int    `gorm:"column:id"`
	EventIDFK int    `gorm:"column:event_id"`
	BuyerIDFK int    `gorm:"column:buyer_id"`
	Status    string `gorm:"column:status;default:'valid';check:status IN ('valid', 'used', 'expired')"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Event     Event
}

func (Ticket) TableName() string {
	return "ticket"
}
