package domain

import (
    "time"
)

type HistoryItem struct {
	Id       int `gorm:"column:id"`
	HistoryIDFK int `gorm:"column:history_id"`
	EventIDFK int `gorm:"column:event_id"`
	Price float64 `gorm:"column:price"`
	Qty int `gorm:"column:qty"`
	Subtotal float64 `gorm:"column:subtotal"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}