package seller_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
)

type EventService interface {
	GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error)
	ViewEvent(eventId int) (buyer_entity.EventEntity, error)
	SaveEvents(userID int, request seller_web.CreateEventRequest) (seller_entity.EventEntity, error)
	GetEventByID(id int) (domain.Event, error)
	DeleteEvent(eventId int) error
	UpdateEvent(request seller_web.UpdateEventRequest, pathId int) (seller_entity.EventEntity, error)
	UpdateEventStatus(request seller_web.UpdateEventStatusRequest, id int) (seller_entity.EventEntity, error)
	CheckInTicket(eventID int, ticketID int) error
}
