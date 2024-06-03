package eventrepository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (repository *EventRepositoryImpl) CreateEvent(event domain.Event) (domain.Event, error) {
	err := repository.db.Create(&event).Error
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}
func (repository *EventRepositoryImpl) GetEventByID(id int) (domain.Event, error) {
	var event domain.Event
	err := repository.db.First(&event, id).Error
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

//update
func (repository *EventRepositoryImpl) UpdateEvent(event domain.Event) (domain.Event, error) {
	err := repository.db.Model(domain.Event{}).Where("id = ?", event.ID).Updates(event).Error
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}



