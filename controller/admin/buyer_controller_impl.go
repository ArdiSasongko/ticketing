package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
)

type BuyerControllerImpl struct {
	buyerService admin_service.BuyerService
}

func NewBuyerController(buyerService admin_service.BuyerService) *BuyerControllerImpl {
	return &BuyerControllerImpl{
		buyerService: buyerService,
	}
}

// List godoc
// @Summary Buyer (List)
// @Description Buyer (List)
// @Tags [Admin] Buyer
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param filters query string false "Filters"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/buyers [get]
func (controller *BuyerControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	result, err := controller.buyerService.GetBuyers(filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

// View godoc
// @Summary Buyer (View)
// @Description Buyer (View)
// @Tags [Admin] Buyer
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/buyers/{id} [get]
func (controller *BuyerControllerImpl) View(c echo.Context) error {
	buyerID, _ := strconv.Atoi(c.Param("id"))

	result, err := controller.buyerService.GetBuyerByID(buyerID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
