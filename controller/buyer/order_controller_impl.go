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

// List godoc
// @Summary Order (List)
// @Description Order (List)
// @Tags [Buyer] Order
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param filter[number] query string false "Filter"
// @Param filter[payment_status] query string false "Filter"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /buyer/orders [get]
func (controller *OrderControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	histories, err := controller.orderService.ListOrder(filters, sort, limit, page)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", histories))
}

// View godoc
// @Summary Order (View)
// @Description Order (View)
// @Tags [Buyer] Order
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "id"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /buyer/orders/{id} [get]
func (controller *OrderControllerImpl) View(c echo.Context) error {
	historyId, _ := strconv.Atoi(c.Param("id"))

	history, err := controller.orderService.ViewOrder(historyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success", history))
}

// Create godoc
// @Summary Order (Create)
// @Description Order (Create)
// @Tags [Buyer] Order
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param event body buyer_web.CreateOrderRequest true "Create Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/orders [post]
func (controller *OrderControllerImpl) Create(c echo.Context) error {
	authId, err := helper.GetAuthId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	input := new(buyer_web.CreateOrderRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	if err := c.Validate(input); err != nil {
		return err
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
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, createOrderErr.Error(), nil))
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, model.ResponseToClient(http.StatusCreated, "Create Order Success", data))
}

// Pay godoc
// @Summary Order (Pay)
// @Description Order (Pay)
// @Tags [Buyer] Order
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "Event ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/orders/{id}/pay [patch]
func (controller *OrderControllerImpl) Pay(c echo.Context) error {
	orderId, _ := strconv.Atoi(c.Param("id"))

	_, err := controller.orderService.PayOrder(orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Pay Order Success", nil))
}

// Delete godoc
// @Summary Order (Delete)
// @Description Order (Delete)
// @Tags [Buyer] Order
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/orders/{id} [delete]
func (controller *OrderControllerImpl) Delete(c echo.Context) error {
	orderId, _ := strconv.Atoi(c.Param("id"))

	if err := controller.orderService.DeleteOrder(orderId); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Delete Order Success", nil))
}
