package history_repository

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type HistoryRepo interface {
	GetHistory(userID int) ([]domain.History, error)
}
