package buyer_service

import (
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
)

type EventServiceImpl struct {
	repository buyer_repository.EventRepository
}

func NewEventService(repository buyer_repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) GetEventList(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error) {
	events, err := service.repository.ListEvents(filters, sort, limit, page)
	if err != nil {
		return []buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventListEntity(events), nil
}
