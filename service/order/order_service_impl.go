package buyer_service

import (
	"strconv"
	"strings"
	"time"

	"github.com/ArdiSasongko/ticketing_app/helper"
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	buyer_web "github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	order_repository "github.com/ArdiSasongko/ticketing_app/repository/order"
)

type Order struct {
	Repo  order_repository.OrderRepositoryInterface
	Token helper.TokenUseCase
}

func (service *Order) Order(id int, req buyer_web.Order) (map[string]interface{}, error) {
	order, errOrder := service.Repo.GetByID(id)

	if errOrder != nil {
		return nil, errOrder
	}

	if req.Qty != 0 {
		order.Qty = req.Qty
	}

	result, errUpdate := service.Repo.Order(id, order)

	if errUpdate != nil {
		return nil, errUpdate
	}

	data := helper.CustomResponse{
		"qty": result.Qty,
	}

	return data, nil
}

func (service *Order) GenerateHistoryItem() error {
	sources, err := service.Repo.CopyEvent()
	if err != nil {
		return err
	}

	var destinations []buyer_entity.HistoryItemEntity
	for _, source := range sources {
		destinations = append(destinations, buyer_entity.HistoryItemEntity{
			EventIDFK: source.ID,
			Price:     source.Price,
		})
	}

	return service.Repo.PasteHistory(destinations)
}

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0

	for num > 0 {
		// calculate the number of times this num is completly divisible by values[i]
		// times will only be > 0, when num >= values[i]
		k := num / values[i]
		for j := 0; j < k; j++ {
			// buildup roman numeral
			roman += symbols[i]

			// reduce the value of num.
			num -= values[i]
		}
		i++
	}
	return roman
}

func (service *Order) GenerateOrderNumber() (string, error) {
	now := time.Now()
	month := now.Month()
	year := strconv.Itoa(now.Year())

	dateCode := "/ORDER/" + intToRoman(int(month)) + "/" + year

	latestOrder, err := service.Repo.GetLatestOrder()
	if err != nil {
		latestOrder = &buyer_entity.HistoryEntity{Number: "0"}
	}

	latestNumber, _ := strconv.Atoi(strings.Split(latestOrder.Number, "/")[0])
	number := latestNumber + 1

	invoiceNumber := strconv.Itoa(number) + dateCode

	return invoiceNumber, nil
}

func (service *Order) PayOrder(BuyyerIDFK uint) (*buyer_entity.HistoryEntity, error) {
	order, err := service.Repo.GetOrderByID(BuyyerIDFK)
	if err != nil {
		return nil, err
	}

	if order.PaymentStatus == buyer_entity.HistoryEntity.OrderPaymentStatusPaid {
		return order, nil // Order already paid, nothing to update
	}

	order.PaymentStatus = buyer_entity.HistoryEntity.OrderPaymentStatusPaid
	order.Status = 1

	if err := service.Repo.UpdateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}
