package buyer_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	"gorm.io/gorm"
)

type TicketRepositoryImpl struct {
	ticketQueryBuilder buyer_query_builder.TicketQueryBuilder
	db                 *gorm.DB
}

func NewTicketRepository(ticketQueryBuilder buyer_query_builder.TicketQueryBuilder, db *gorm.DB) *TicketRepositoryImpl {
	return &TicketRepositoryImpl{
		ticketQueryBuilder: ticketQueryBuilder,
		db:                 db,
	}
}

func (repo *TicketRepositoryImpl) ListTickets(filters map[string]string, sort string, limit int, page int) ([]domain.Ticket, error) {
	var tickets []domain.Ticket

	ticketQueryBuilder, err := repo.ticketQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := ticketQueryBuilder.Preload("Event.Seller").Find(&tickets).Error
	if err1 != nil {
		return []domain.Ticket{}, err1
	}
	return tickets, nil
}

func (repo *TicketRepositoryImpl) GetTicket(ticketId int) (domain.Ticket, error) {
	var ticket domain.Ticket

	if getTicketErr := repo.db.Preload("Event.Seller").Where("id = ?", ticketId).First(&ticket).Error; getTicketErr != nil {
		return domain.Ticket{}, getTicketErr
	}

	return ticket, nil
}
