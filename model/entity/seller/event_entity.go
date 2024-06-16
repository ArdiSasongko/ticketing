package seller_entity

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"time"
)

type EventEntity struct {
	EventID  int       `json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
	Qty      int       `json:"qty"`
	Category string    `json:"category"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
}

func ToEventEntity(event domain.Event) EventEntity {
	return EventEntity{
		EventID:  event.ID,
		Name:     event.Name,
		Date:     event.Date,
		Location: event.Location,
		Qty:      event.Qty,
		Category: event.Category,
		Price:    event.Price,
		Status:   event.Status,
	}
}

func ToEventListEntity(events []domain.Event) []EventEntity {
	var eventData []EventEntity

	for _, event := range events {
		eventData = append(eventData, ToEventEntity(event))

	}
	return eventData
}
