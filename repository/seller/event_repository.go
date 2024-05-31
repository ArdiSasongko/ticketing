package seller

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type EventRepository interface {
	ListEvents(sellerId int, filters map[string]string, sort string, limit int, page int) ([]domain.Events, error)
}
