package history_repository

import (
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"gorm.io/gorm"
)

type HistoryRepoImpl struct {
	DB *gorm.DB
}

func NewHistoryRepoImpl(db *gorm.DB) *HistoryRepoImpl {
	return &HistoryRepoImpl{
		DB: db,
	}
}

func (repo *HistoryRepoImpl) GetHistory(userID int) ([]domain.History, error) {
	var history []domain.History
	err := repo.DB.Where("buyer_id = ?", userID).Find(&history).Error

	if err != nil {
		return nil, err
	}

	return history, nil
}
