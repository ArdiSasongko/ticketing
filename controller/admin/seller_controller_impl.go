package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
)

type SellerControllerImpl struct {
	sellerService admin_service.SellerService
}

func NewSellerController(sellerService admin_service.SellerService) *SellerControllerImpl {
	return &SellerControllerImpl{
		sellerService: sellerService,
	}
}

// List godoc
// @Summary Seller (List)
// @Description Seller (List)
// @Tags [Admin] Seller
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param filters query string false "Filters"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/sellers [get]
func (controller *SellerControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	result, err := controller.sellerService.GetSellers(filters, sort, limit, page)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

// View godoc
// @Summary Seller (View)
// @Description Seller (View)
// @Tags [Admin] Seller
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/sellers/{id} [get]
func (controller *SellerControllerImpl) View(c echo.Context) error {
	sellerID, _ := strconv.Atoi(c.Param("id"))

	result, err := controller.sellerService.GetSellerByID(sellerID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
