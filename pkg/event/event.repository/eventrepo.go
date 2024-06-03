package eventrepository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type EventRepository interface {
	CreateEvent(event domain.Event) (domain.Event, error)
	GetEventByID(id int) (domain.Event, error)
	UpdateEvent(user domain.Event) (domain.Event, error)
}
