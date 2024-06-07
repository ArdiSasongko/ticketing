package event_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type EventRepository interface {
	GetEvent(Id int) (domain.Event, error)
	GetEvents() ([]domain.Event, error)
}
