package buyer_repository

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	buyer_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	"gorm.io/gorm"
	"time"
)

type OrderRepositoryImpl struct {
	orderQueryBuilder buyer_query_builder.OrderQueryBuilder
	DB                *gorm.DB
}

func NewOrderRepository(orderQueryBuilder buyer_query_builder.OrderQueryBuilder, db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		orderQueryBuilder: orderQueryBuilder,
		DB:                db,
	}
}

func (repo *OrderRepositoryImpl) WithTx(orderQueryBuilder buyer_query_builder.OrderQueryBuilder, tx *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{orderQueryBuilder, tx}
}

func (repo *OrderRepositoryImpl) ListHistory(filters map[string]string, sort string, limit int, page int) ([]domain.History, error) {
	var histories []domain.History

	orderQueryBuilder, err := repo.orderQueryBuilder.GetBuilder(filters, sort, limit, page)
	if err != nil {
		return nil, err
	}

	err1 := orderQueryBuilder.Find(&histories).Error
	if err1 != nil {
		return []domain.History{}, err1
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

func (repo *OrderRepositoryImpl) GetLatestOrder(buyerId int) (domain.History, error) {
	var latestOrder domain.History

	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)
	err := repo.DB.First(&latestOrder).Where("created_at >= ?", today).Where("created_at < ?", tomorrow).Where("buyer_id = ?", buyerId).Order("created_at DESC").Error
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
