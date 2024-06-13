package buyer_controller

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type OrderControllerImpl struct {
	orderService buyer_service.OrderService
}

func NewOrderController(orderService buyer_service.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{
		orderService: orderService,
	}
}

func (controller *OrderControllerImpl) ListOrder(c echo.Context) error {
	histories, err := controller.orderService.ListOrder()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "list order success", histories))
}

func (controller *OrderControllerImpl) ViewOrder(c echo.Context) error {
	historyId, _ := strconv.Atoi(c.Param("id"))

	history, err := controller.orderService.ViewOrder(historyId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "view order success", history))
}

func (controller *OrderControllerImpl) CreateOrder(c echo.Context) error {
	authId, err := helper.GetAuthId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	input := new(buyer_web.OrderRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	tx := app.DBConnection().Begin()
	orderService := controller.orderService.WithTx(tx)

	if deleteActiveOrderErr := orderService.DeleteActiveOrder(authId); deleteActiveOrderErr != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	data, createOrderErr := orderService.CreateOrder(*input, authId)
	if createOrderErr != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, model.ResponseToClient(http.StatusCreated, "Create Order Success", data))
}

func (controller *OrderControllerImpl) PayOrder(c echo.Context) error {
	orderId, _ := strconv.Atoi(c.Param("id"))

	_, err := controller.orderService.PayOrder(orderId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Pay Order Success", nil))
}

func (controller *OrderControllerImpl) DeleteOrder(c echo.Context) error {
	orderId, _ := strconv.Atoi(c.Param("id"))

	if err := controller.orderService.DeleteOrder(orderId); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Delete Order Success", nil))
}
