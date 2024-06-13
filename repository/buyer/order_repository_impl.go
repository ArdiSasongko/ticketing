package buyer_repository

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"gorm.io/gorm"
	"time"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func (repo *OrderRepositoryImpl) WithTx(tx *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{tx}
}

func (repo *OrderRepositoryImpl) ListHistory() ([]domain.History, error) {
	var histories []domain.History

	if err := repo.DB.Preload("HistoryItems.Event").Preload("Buyer").Find(&histories).Error; err != nil {
		return []domain.History{}, err
	}

	return histories, nil
}

func (repo *OrderRepositoryImpl) CreateHistory(history domain.History) (domain.History, error) {
	if err := repo.DB.Preload("Buyer").Preload("HistoryItems").Create(&history).Error; err != nil {
		return domain.History{}, err
	}

	return history, nil
}

func (repo *OrderRepositoryImpl) UpdateHistory(history domain.History) (domain.History, error) {
	if err := repo.DB.Updates(&history).Error; err != nil {
		return domain.History{}, err
	}

	return history, nil
}

func (repo *OrderRepositoryImpl) CreateHistoryItem(historyItem domain.HistoryItem) (domain.HistoryItem, error) {
	if err := repo.DB.Create(&historyItem).Error; err != nil {
		return domain.HistoryItem{}, err
	}

	return historyItem, nil
}

func (repo *OrderRepositoryImpl) GetEvent(id int) (domain.Event, error) {
	var event domain.Event
	repo.DB.Where("id = ?", id).First(&event)

	return event, nil
}

func (repo *OrderRepositoryImpl) UpdateEvent(event domain.Event) (domain.Event, error) {
	repo.DB.Save(&event)

	return event, nil
}

func (repo *OrderRepositoryImpl) GetLatestOrder() (domain.History, error) {
	var latestOrder domain.History

	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)
	err := repo.DB.Where("created_at >= ?", today).Where("created_at < ?", tomorrow).Order("created_at DESC").First(&latestOrder).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.History{}, nil
		}
		return domain.History{}, err
	}
	return latestOrder, nil
}

func (repo *OrderRepositoryImpl) GetHistory(historyId int) (domain.History, error) {
	var history domain.History

	err := repo.DB.Preload("HistoryItems.Event.Seller").First(&history, "id = ?", historyId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.History{}, errors.New("order not found")
		}
		return domain.History{}, err
	}

	return history, nil
}

func (repo *OrderRepositoryImpl) GetActiveHistory(buyerId int) (domain.History, error) {
	var history domain.History
	if err := repo.DB.Model(&domain.History{}).Preload("HistoryItems").Where("payment_status = ?", enum.PaymentStatusPending).Where("buyer_id = ?", buyerId).Find(&history).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.History{}, nil
		}
		return domain.History{}, err
	}
	print(history.Number)

	return history, nil
}

func (repo *OrderRepositoryImpl) DeleteHistory(history domain.History) error {
	if err := repo.DB.Delete(&history).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OrderRepositoryImpl) DeleteHistoryItem(historyItem domain.HistoryItem) error {
	return repo.DB.Delete(&historyItem).Error
}
