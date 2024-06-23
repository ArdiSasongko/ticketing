package buyer_service

import "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"

type TicketService interface {
	GetTicketList(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.TicketEntity, error)
	ViewTicket(ticketId int) (buyer_entity.TicketEntity, error)
}
