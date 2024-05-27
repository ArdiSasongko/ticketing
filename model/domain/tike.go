package domain

import (
    "time"
)

type tiket struct {
	Id       int `gorm:"column:id"`
	EventIDFK int `gorm:"column:event_id"`
	BuyyerIDFK int `gorm:"column:buyer_id"`
	Date        time.Time `gorm:"column:date"` 
	Location   string `gorm:"column:location"`
	Qty        int `gorm:"column:qty"`
	Price      float64   `gorm:"column:price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}