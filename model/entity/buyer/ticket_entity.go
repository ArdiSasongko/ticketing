package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type TicketEntity struct {
	Id        int         `json:"id"`
	BuyerIDFK int         `json:"buyer_id"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Event     EventEntity `json:"event"`
}

func ToTicketEntity(ticket domain.Ticket) TicketEntity {
	return TicketEntity{
		Id:        ticket.Id,
		BuyerIDFK: ticket.BuyerIDFK,
		Status:    ticket.Status,
		CreatedAt: ticket.CreatedAt,
		UpdatedAt: ticket.UpdatedAt,
		Event:     ToEventEntity(ticket.Event),
	}
}

func ToTicketListEntity(tickets []domain.Ticket) []TicketEntity {
	var ticketList []TicketEntity

	for _, ticket := range tickets {
		ticketList = append(ticketList, ToTicketEntity(ticket))
	}

	return ticketList
}
