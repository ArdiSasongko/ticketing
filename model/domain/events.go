package domain

import "time"

type Events struct {
	EventID   int     `gorm:"column:id;primaryKey;autoIncrement"`
	Seller_ID int     `gorm:"column:seller_id`
	Name      string  `gorm:"column:name"`
	Date      string  `gorm:"column:date"`
	Location  string  `gorm:"column:location"`
	Qty       int     `gorm:"column:qty"`
	Category  string  `gorm:"column:category"`
	Price     float32 `gorm:"column:price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
