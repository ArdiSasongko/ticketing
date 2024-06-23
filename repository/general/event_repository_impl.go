package general_repository

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	general_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/general"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	eventQueryBuilder general_query_builder.EventQueryBuilder
	db                *gorm.DB
}

func NewEventRepository(eventQueryBuilder general_query_builder.EventQueryBuilder, db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		eventQueryBuilder: eventQueryBuilder,
		db:                db,
	}
}

func (repo *EventRepositoryImpl) List(filters map[string]string, sort string, limit int, page int) ([]domain.Event, error) {
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

func (repo *EventRepositoryImpl) View(eventId int) (domain.Event, error) {
	var event domain.Event

	if err := repo.db.Where("id = ?", eventId).Take(&event).Error; err != nil {
		return domain.Event{}, errors.New("event not found")
	}

	return event, nil
}
