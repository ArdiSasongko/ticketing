package event_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/event"
	"github.com/ArdiSasongko/ticketing_app/model/web/event"
)

type EventService interface {
	GetEventList() ([]event_entity.EventEntity, error)
	DeleteEvent(request event_web.EventDeleteServiceRequest, pathId int) (map[string]interface{}, error)
}
