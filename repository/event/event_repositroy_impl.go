package repository

import "fmt"

type EventRepositoryImpl struct {
	db *gorm.DB
}

func (repo *EventRepositoryImpl) GetUsers() ([]domain.events, error) {
	var events []domain.events

	err := repo.db.Find(&events).Error

	if err != nil {
		return []domain.events{}, err
	}
	fmt.Println(events)
	return events, nil
}

func (repo *EventRepositoryImpl) DeleteEvent(event domain.Events) (domain.Events, error) {
	err := repo.db.Model(domain.Events{}).Where("id=?", user.EventID).Delete(event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}
