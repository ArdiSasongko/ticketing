package general_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/event"
	"github.com/ArdiSasongko/ticketing_app/repository/general"
)

type EventServiceImpl struct {
	eventRepo general_repository.EventRepository
}

func NewEventService(eventRepo general_repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepo: eventRepo,
	}
}

func (service *EventServiceImpl) GetEventList(filters map[string]string, sort string, limit int, page int) ([]event_entity.EventEntity, error) {
	getEventList, errGetEventList := service.eventRepo.List(filters, sort, limit, page)

	if errGetEventList != nil {
		return []event_entity.EventEntity{}, errGetEventList
	}

	return event_entity.ToEventListEntity(getEventList), nil
}

func (service *EventServiceImpl) GetEvent(eventId int) (event_entity.EventEntity, error) {
	getEventList, errGetEventList := service.eventRepo.View(eventId)

	if errGetEventList != nil {
		return event_entity.EventEntity{}, errGetEventList
	}

	return event_entity.ToEventEntity(getEventList), nil
}
