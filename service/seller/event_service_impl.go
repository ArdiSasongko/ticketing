package seller_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
)

type EventServiceImpl struct {
	repository seller_repository.EventRepository
}

func NewEventService(repository seller_repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error) { // todo
	events, err := service.repository.ListEvents(sellerId, filters, sort, limit, page)
	if err != nil {
		return []buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventListEntity(events), nil
}
