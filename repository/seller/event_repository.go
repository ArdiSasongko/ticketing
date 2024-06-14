package seller_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type EventRepository interface {
	ListEvents(sellerId int, filters map[string]string, sort string, limit int, page int) ([]domain.Event, error)
	CreateEvent(event domain.Event) (domain.Event, error)
	GetEventByID(id int) (domain.Event, error)
	UpdateEvent(user domain.Event) (domain.Event, error)
	CheckInTicket(eventID int, ticketID int) error
}
