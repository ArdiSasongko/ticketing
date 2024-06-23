package buyer_service

import (
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
)

type TicketServiceImpl struct {
	ticketRepository buyer_repository.TicketRepository
}

func NewTicketService(ticketRepository buyer_repository.TicketRepository) *TicketServiceImpl {
	return &TicketServiceImpl{
		ticketRepository: ticketRepository,
	}
}

func (service *TicketServiceImpl) GetTicketList(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.TicketEntity, error) {
	events, err := service.ticketRepository.ListTickets(filters, sort, limit, page)
	if err != nil {
		return []buyer_entity.TicketEntity{}, err
	}

	return buyer_entity.ToTicketListEntity(events), nil
}

func (service *TicketServiceImpl) ViewTicket(ticketId int) (buyer_entity.TicketEntity, error) {
	ticket, getTicketErr := service.ticketRepository.GetTicket(ticketId)
	if getTicketErr != nil {
		return buyer_entity.TicketEntity{}, getTicketErr
	}

	return buyer_entity.ToTicketEntity(ticket), nil
}
