package general_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type EventRepository interface {
	List(filters map[string]string, sort string, limit int, page int) ([]domain.Event, error)
	View(eventId int) (domain.Event, error)
}
