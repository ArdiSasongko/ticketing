package buyer_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type TicketRepository interface {
	ListTickets(filters map[string]string, sort string, limit int, page int) ([]domain.Ticket, error)
	GetTicket(ticketId int) (domain.Ticket, error)
}
