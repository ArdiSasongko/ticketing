package buyer_controller

import (
	"net/http"

	"github.com/ArdiSasongko/ticketing_app/helper"
	service "github.com/ArdiSasongko/ticketing_app/service/order"
	"github.com/labstack/echo/v4"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
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
