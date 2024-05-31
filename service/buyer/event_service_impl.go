package buyer

import (
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
)

type EventServiceImpl struct {
	repository buyer.EventRepository
}

func NewEventService(repository buyer.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) GetEventList(filters map[string]string, sort string, limit int, page int) ([]entity.EventEntity, error) {
	events, err := service.repository.ListEvents(filters, sort, limit, page)
	if err != nil {
		return []entity.EventEntity{}, err
	}

	return entity.ToEventListEntity(events), nil
}
