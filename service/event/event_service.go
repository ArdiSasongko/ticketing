package event_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/event"
)

type EventService interface {
	GetEventList() ([]event_entity.EventEntity, error)
}
