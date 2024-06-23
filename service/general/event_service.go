package general_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/event"
)

type EventService interface {
	GetEventList(filters map[string]string, sort string, limit int, page int) ([]event_entity.EventEntity, error)
	GetEvent(eventId int) (event_entity.EventEntity, error)
}
