package buyer_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
)

type EventRepositoryImpl struct {
	eventQueryBuilder buyer_query_builder.EventQueryBuilder
}

func NewEventRepository(eventQueryBuilder buyer_query_builder.EventQueryBuilder) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		eventQueryBuilder: eventQueryBuilder,
	}
}

func (repo *EventRepositoryImpl) ListEvents(filters map[string]string, sort string, limit int, page int) ([]domain.Event, error) {
	var events []domain.Event

	eventQueryBuilder, err := repo.eventQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := eventQueryBuilder.Find(&events).Error
	if err1 != nil {
		return []domain.Event{}, err1
	}
	return events, nil
}
