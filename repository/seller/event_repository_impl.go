package seller_repository

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/query_builder/seller"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	eventQueryBuilder seller_query_builder.EventQueryBuilder
	db                *gorm.DB
}

func NewEventRepository(eventQueryBuilder seller_query_builder.EventQueryBuilder) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		eventQueryBuilder: eventQueryBuilder,
		db:                app.DBConnection(),
	}
}

func (repo *EventRepositoryImpl) ListEvents(sellerId int, filters map[string]string, sort string, limit int, page int) ([]domain.Event, error) {
	var events []domain.Event

	eventQueryBuilder, err := repo.eventQueryBuilder.GetBuilder(sellerId, filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := eventQueryBuilder.Find(&events).Error
	if err1 != nil {
		return []domain.Event{}, err1
	}
	return events, nil
}

func (repo *EventRepositoryImpl) CreateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Create(&event).Error
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}
func (repo *EventRepositoryImpl) GetEventByID(id int) (domain.Event, error) {
	var event domain.Event
	err := repo.db.First(&event, id).Error
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (repo *EventRepositoryImpl) UpdateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Model(domain.Event{}).Where("id = ?", event.ID).Updates(event).Error
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

func (repo *EventRepositoryImpl) CheckInTicket(eventID int, ticketID int) error {

	var ticket domain.Ticket
	if err := repo.db.Where("id = ?", ticketID).First(&ticket).Error; err != nil {
		return err
	}

	if ticket.Status == "Used" {
		return errors.New("Ticket has already been checked in")
	}

	ticket.Status = "Used"
	if err := repo.db.Save(&ticket).Error; err != nil {
		return err
	}

	return nil
}

func (repo *EventRepositoryImpl) DeleteEventById(eventId int) error {
	return repo.db.Delete(&domain.Event{}, eventId).Error
}
