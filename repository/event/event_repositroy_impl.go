package event_repository

import (
	"fmt"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func (repo *EventRepositoryImpl) GetEvents() ([]domain.Event, error) {
	var events []domain.Event

	err := repo.db.Find(&events).Error

	if err != nil {
		return []domain.Event{}, err
	}
	fmt.Println(events)
	return events, nil
}

func (repo *EventRepositoryImpl) DeleteEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Model(domain.Event{}).Where("id=?", event.ID).Delete(event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}
