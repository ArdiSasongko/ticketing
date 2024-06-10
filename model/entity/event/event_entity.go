package event_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type EventEntity struct {
	EventID   int       `json:"event_id"`
	Seller_ID int       `json:"seller_id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Location  string    `json:"location"`
	Qty       int       `json:"qty"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
}

func ToEventEntity(id int, name string, seller_id int, date time.Time, location string, qty int, category string, price float64) EventEntity {
	return EventEntity{
		EventID:   id,
		Seller_ID: seller_id,
		Name:      name,
		Date:      date,
		Location:  location,
		Qty:       qty,
		Category:  category,
		Price:     price,
	}
}
func ToEventListEntity(events []domain.Event) []EventEntity {
	eventData := []EventEntity{}

	for _, event := range events {
		eventData = append(eventData, ToEventEntity(event.ID, event.Name, event.SellerID, event.Date, event.Location, event.Qty, event.Category, event.Price))

	}
	return eventData
}
