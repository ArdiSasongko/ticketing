package seller_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
)

type EventService interface {
	GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error)
	ViewEvent(eventId int) (buyer_entity.EventEntity, error)
	SaveEvents(userID int, request seller_web.CreateEventsRequest) (map[string]interface{}, error)
	GetEventByID(id int) (domain.Event, error)
	DeleteEvent(eventId int) error
	UpdateEvent(request seller_web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	CheckInTicket(eventID int, ticketID int) error
}
