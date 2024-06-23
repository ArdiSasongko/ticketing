package domain

import "time"

type Event struct {
	EventID   int       `gorm:"column:id;primaryKey;autoIncrement"` // todo
	SellerID  int       `gorm:"column:seller_id"`
	Name      string    `gorm:"column:name"`
	Date      time.Time `gorm:"column:date"`
	Location  string    `gorm:"column:location"`
	Qty       int       `gorm:"column:qty"`
	Category  string    `gorm:"column:category"`
	Price     float64   `gorm:"column:price"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Seller    Sellers `gorm:"foreignKey:seller_id"`
}

func (Event) TableName() string {
	return "event"
}
