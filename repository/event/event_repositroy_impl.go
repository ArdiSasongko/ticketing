package repository

import (
	"fmt"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func (repo *EventRepositoryImpl) GetEvents() ([]domain.Events, error) {
	var events []domain.Events

	err := repo.db.Find(&events).Error

	if err != nil {
		return []domain.Events{}, err
	}
	fmt.Println(events)
	return events, nil
}

func (repo *EventRepositoryImpl) DeleteEvent(event domain.Events) (domain.Events, error) {
	err := repo.db.Model(domain.Events{}).Where("id=?", event.EventID).Delete(event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}
