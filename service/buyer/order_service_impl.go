package buyer_service

import (
	"errors"
	"fmt"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type OrderServiceImpl struct {
	db               *gorm.DB
	orderRepository  buyer_repository.OrderRepository
	sellerRepository buyer_repository.SellerRepository
}

func NewOrderService(
	db *gorm.DB,
	orderRepository buyer_repository.OrderRepository,
	sellerRepository buyer_repository.SellerRepository,
) *OrderServiceImpl {
	return &OrderServiceImpl{
		db:               db,
		orderRepository:  orderRepository,
		sellerRepository: sellerRepository,
	}
}

func (service *OrderServiceImpl) WithTx(tx *gorm.DB) OrderService {
	return &OrderServiceImpl{tx, service.orderRepository, service.sellerRepository}
}

func (service *OrderServiceImpl) ListOrder(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.HistoryLiteEntity, error) {
	histories, err := service.orderRepository.ListHistory(filters, sort, limit, page)

	if err != nil {
		return []buyer_entity.HistoryLiteEntity{}, err
	}

	return buyer_entity.ToHistoryLiteEntityCollection(histories), nil
}

func (service *OrderServiceImpl) ViewOrder(historyId int) (buyer_entity.HistoryEntity, error) {
	history, err := service.orderRepository.GetHistory(historyId)

	if err != nil {
		return buyer_entity.HistoryEntity{}, err
	}

	return buyer_entity.ToHistoryEntity(history), nil
}

func (service *OrderServiceImpl) CreateOrder(request buyer_web.CreateOrderRequest, buyerId int) (buyer_entity.HistoryEntity, error) {
	number, generateOrderNumberErr := service.generateOrderNumber(buyerId)
	if generateOrderNumberErr != nil {
		return buyer_entity.HistoryEntity{}, generateOrderNumberErr
	}

	history := domain.History{}
	history.BuyerIDFK = buyerId
	history.Number = number
	history.PaymentStatus = string(enum.PaymentStatusPending)
	history.Total = 0
	history, createHistoryErr := service.orderRepository.CreateHistory(history)
	if createHistoryErr != nil {
		return buyer_entity.HistoryEntity{}, createHistoryErr
	}

	event, getEventErr := service.orderRepository.GetEvent(request.EventID)
	if getEventErr != nil {
		return buyer_entity.HistoryEntity{}, getEventErr
	}

	if event.Status != string(enum.EventStatusActive) {
		return buyer_entity.HistoryEntity{}, errors.New("event is not active")
	}
	if event.Qty < request.Qty {
		return buyer_entity.HistoryEntity{}, errors.New("not enough ticket")
	}

	historyItem := domain.HistoryItem{}
	historyItem.HistoryIDFK = history.Id
	historyItem.EventIDFK = event.EventID
	historyItem.Price = event.Price
	historyItem.Qty = request.Qty
	historyItem.Subtotal = historyItem.Price * float64(historyItem.Qty)
	historyItem, createHistoryItemErr := service.orderRepository.CreateHistoryItem(historyItem)
	if createHistoryItemErr != nil {
		return buyer_entity.HistoryEntity{}, createHistoryItemErr
	}

	history.Total += historyItem.Subtotal
	history, updateHistoryErr := service.orderRepository.UpdateHistory(history)
	if updateHistoryErr != nil {
		return buyer_entity.HistoryEntity{}, updateHistoryErr
	}
	history, getHistoryErr := service.orderRepository.GetHistory(history.Id)
	if getHistoryErr != nil {
		return buyer_entity.HistoryEntity{}, getHistoryErr
	}

	return buyer_entity.ToHistoryEntity(history), nil
}

func (service *OrderServiceImpl) DeleteActiveOrder(buyerId int) error {
	activeHistory, getActiveHistoryErr := service.orderRepository.GetActiveHistory(buyerId)
	if getActiveHistoryErr != nil {
		return getActiveHistoryErr
	}

	if activeHistory.Id != 0 {
		var historyItems []domain.HistoryItem
		historyItems = activeHistory.HistoryItems
		for _, historyItem := range historyItems {
			if deleteHistoryItemErr := service.orderRepository.DeleteHistoryItem(historyItem); deleteHistoryItemErr != nil {
				return deleteHistoryItemErr
			}
		}

		if deleteHistoryErr := service.orderRepository.DeleteHistory(activeHistory); deleteHistoryErr != nil {
			return deleteHistoryErr
		}
	}

	return nil
}

func (service *OrderServiceImpl) generateOrderNumber(buyerId int) (string, error) {
	yearInt, monthInt, dayInt := time.Now().Date()
	year := strconv.Itoa(yearInt)
	month := strconv.Itoa(int(monthInt))
	day := strconv.Itoa(dayInt)
	if len(month) == 1 {
		month = fmt.Sprintf("0%s", month)
	}
	if len(day) == 1 {
		day = fmt.Sprintf("0%s", day)
	}
	latestOrder, err := service.orderRepository.GetLatestOrder(buyerId)
	if err != nil {
		return "", err
	}
	var codeInt int
	if latestOrder.Id != 0 {
		codeStr := latestOrder.Number[11:]
		codeInt, _ = strconv.Atoi(codeStr)
	} else {
		codeInt = 0
	}
	codeStr := strconv.Itoa(codeInt + 1)
	codeStr = fmt.Sprintf("%s%s", strings.Repeat("0", 4-len(codeStr)), codeStr)

	return fmt.Sprintf("INV%s%s%s%s", year, month, day, codeStr), nil
}

func (service *OrderServiceImpl) PayOrder(orderId int) (buyer_entity.HistoryEntity, error) {
	history, err := service.orderRepository.GetHistory(orderId)
	if err != nil {
		return buyer_entity.HistoryEntity{}, err
	}

	if history.PaymentStatus == string(enum.PaymentStatusPaid) {
		return buyer_entity.HistoryEntity{}, errors.New("order already paid")
	}

	for _, historyItem := range history.HistoryItems {
		event := historyItem.Event
		event.Qty -= historyItem.Qty
		if event.Qty < 0 {
			return buyer_entity.HistoryEntity{}, errors.New("not enough ticket")
		}
		event, err = service.orderRepository.UpdateEvent(event)
		if err != nil {
			return buyer_entity.HistoryEntity{}, err
		}

		seller := event.Seller
		seller, updateSellerErr := service.sellerRepository.AddSellerBalance(seller, historyItem.Subtotal)
		if updateSellerErr != nil {
			return buyer_entity.HistoryEntity{}, updateSellerErr
		}

		for i := 0; i < historyItem.Qty; i++ {
			ticket := domain.Ticket{}
			ticket.EventIDFK = event.EventID
			ticket.BuyerIDFK = history.BuyerIDFK
			//ticket.Status = string(enum.TicketStatusValid)
			storeTicketErr := service.orderRepository.CreateTicket(ticket)
			if storeTicketErr != nil {
				return buyer_entity.HistoryEntity{}, storeTicketErr
			}
		}
	}

	history.PaymentStatus = string(enum.PaymentStatusPaid)
	now := time.Now()
	history.PaidAt = &now
	history, err = service.orderRepository.UpdateHistory(history)
	if err != nil {
		return buyer_entity.HistoryEntity{}, err
	}

	return buyer_entity.ToHistoryEntity(history), nil
}

func (service *OrderServiceImpl) DeleteOrder(historyId int) error {
	history, getHistoryErr := service.orderRepository.GetHistory(historyId)
	if getHistoryErr != nil {
		return getHistoryErr
	}

	if history.PaymentStatus == string(enum.PaymentStatusPaid) {
		return errors.New("cannot delete paid order")
	}

	for _, historyItem := range history.HistoryItems {
		err := service.orderRepository.DeleteHistoryItem(historyItem)
		if err != nil {
			return err
		}
	}

	return service.orderRepository.DeleteHistory(history)
}
