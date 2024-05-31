package seller

import (
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
)

type EventServiceImpl struct {
	repository seller.EventRepository
}

func NewEventService(repository seller.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]entity.EventEntity, error) {
	events, err := service.repository.ListEvents(sellerId, filters, sort, limit, page)
	if err != nil {
		return []entity.EventEntity{}, err
	}

	return entity.ToEventListEntity(events), nil
}
