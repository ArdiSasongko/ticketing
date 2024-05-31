package buyer

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
)

type EventRepositoryImpl struct {
	eventQueryBuilder buyer.EventQueryBuilder
}

func NewEventRepository(eventQueryBuilder buyer.EventQueryBuilder) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		eventQueryBuilder: eventQueryBuilder,
	}
}

func (repository *EventRepositoryImpl) ListEvents(filters map[string]string, sort string, limit int, page int) ([]domain.Events, error) {
	var events []domain.Events

	eventQueryBuilder, err := repository.eventQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := eventQueryBuilder.Find(&events).Error
	if err1 != nil {
		return []domain.Events{}, err1
	}
	return events, nil
}
