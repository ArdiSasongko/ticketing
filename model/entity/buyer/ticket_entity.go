package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type TicketEntity struct {
	Id        int       `json:"id"`
	EventIDFK int       `json:"event_id"`
	BuyerIDFK int       `json:"buyer_id"`
	Date      time.Time `json:"date"`
	Location  string    `json:"location"`
	Qty       int       `json:"qty"`
	Price     float64   `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToTicketEntity(ticket domain.Ticket) TicketEntity {
	return TicketEntity{
		Id:        ticket.Id,
		BuyerIDFK: ticket.BuyerIDFK,
		Date:      ticket.Date,
		Location:  ticket.Location,
		Qty:       ticket.Qty,
		CreatedAt: ticket.CreatedAt,
		UpdatedAt: ticket.UpdatedAt,
	}
}

//func GenerateTicket(ticket []domain.Ticket) []TicketEntity {
//	ticketData := []TicketEntity{}
//	historyItemData := []HistoryItemEntity{}
//	for _, source := range historyItemData {
//		ticketData = append(ticketData, TicketEntity{
//
//			EventIDFK: source.EventIDFK,
//			Qty:       source.Qty,
//			Price:     source.Subtotal,
//		})
//	}
//	return ticketData
//}
