package buyer_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	buyer_service "github.com/ArdiSasongko/ticketing_app/service/order"
	service "github.com/ArdiSasongko/ticketing_app/service/order"
	"github.com/labstack/echo/v4"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(service buyer_service.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{OrderService: service}
}

func (controller *OrderControllerImpl) Order(c echo.Context) error {

}

func (controller *OrderControllerImpl) GenerateHistoryItem(c echo.Context) {
	errGenHis := controller.OrderService.GenerateHistoryItem()
	if errGenHis != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, errGenHis.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", nil))
}

func (controller *OrderControllerImpl) GenerateOrderNumber(c echo.Context) error {
	orderNumber, err := controller.OrderService.GenerateOrderNumber()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"order_number": orderNumber})
}

func (controller *OrderControllerImpl) PayOrder(c echo.Context) error {
	orderIDParam := c.Param("ID")
	orderID, err := strconv.ParseUint(orderIDParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	order, err := controller.OrderService.PayOrder(uint(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, order)
}
