package buyer_service

import (
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
)

type EventServiceImpl struct {
	eventRepository buyer_repository.EventRepository
}

func NewEventService(eventRepository buyer_repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepository: eventRepository,
	}
}

func (service *EventServiceImpl) GetEventList(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error) {
	events, err := service.eventRepository.ListEvents(filters, sort, limit, page)
	if err != nil {
		return []buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventListEntity(events), nil
}

func (service *EventServiceImpl) ViewEvent(eventId int) (buyer_entity.EventEntity, error) {
	event, getEventErr := service.eventRepository.GetEvent(eventId)
	if getEventErr != nil {
		return buyer_entity.EventEntity{}, getEventErr
	}

	return buyer_entity.ToEventEntity(event), nil
}
