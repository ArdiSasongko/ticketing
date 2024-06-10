package order_repository

import (
	"errors"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"gorm.io/gorm"
)

type Order struct {
	DB *gorm.DB
}

func (repo *Order) Order(id int, qty domain.Event) (domain.Event, error) {
	if err := repo.DB.Where("id = ?", id).Update("qty", gorm.Expr("qty - ?", qty)).Error; err != nil {
		return domain.Event{}, errors.New("maaf, tiket habis")
	}
	return qty, nil
}

func (r *Order) CopyEvent() ([]buyer_entity.EventEntity, error) {
	var sources []buyer_entity.EventEntity
	result := r.DB.Find(&sources)
	return sources, result.Error
}

func (r *Order) PasteHistory(destinations []buyer_entity.HistoryItemEntity) error {
	result := r.DB.Create(&destinations)
	return result.Error
}

func (r *Order) GetLatestOrder() (*buyer_entity.HistoryItemEntity, error) {
	var latestOrder buyer_entity.HistoryItemEntity
	err := r.DB.Debug().Order("created_at DESC").First(&latestOrder).Error
	if err != nil {
		return nil, err
	}
	return &latestOrder, nil
}
