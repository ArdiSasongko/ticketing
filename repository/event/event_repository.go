package event_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type EventRepository interface {
	GetEvent(Id int) (domain.Events, error)
	GetEvents() ([]domain.Events, error)
}
