package buyer_entity

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type EventEntity struct {
	ID        int                        `json:"id"`
	Name      string                     `json:"name"`
	Date      time.Time                  `json:"date"`
	Location  string                     `json:"location"`
	Qty       int                        `json:"qty"`
	Category  string                     `json:"category"`
	Price     float64                    `json:"price"`
	Status    string                     `json:"status"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
	Seller    seller_entity.SellerEntity `json:"seller"`
}

func ToEventEntity(event domain.Event) EventEntity {
	return EventEntity{
		ID:        event.ID,
		Name:      event.Name,
		Date:      event.Date,
		Location:  event.Location,
		Qty:       event.Qty,
		Category:  event.Category,
		Price:     event.Price,
		Status:    event.Status,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
		Seller:    seller_entity.ToSellerEntity(event.Seller),
	}
}

func ToEventListEntity(events []domain.Event) []EventEntity {
	var eventList []EventEntity

	for _, event := range events {
		eventList = append(eventList, ToEventEntity(event))
	}

	return eventList
}
