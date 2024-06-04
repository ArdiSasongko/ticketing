package seller_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	seller_web "github.com/ArdiSasongko/ticketing_app/model/web/seller"
)

type EventService interface {
	GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error) // todo
	SaveEvents(request seller_web.CreateEventsRequest) (map[string]interface{}, error)
	GetEventByID(id int) (domain.Event, error)
	UpdateEvent(request seller_web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
}
