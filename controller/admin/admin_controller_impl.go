package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
)

type AdminControllerImpl struct {
	adminService admin_service.AdminService
}

func NewAdminController(adminService admin_service.AdminService) *AdminControllerImpl {
	return &AdminControllerImpl{
		adminService: adminService,
	}
}

// List godoc
// @Summary Admin (List)
// @Description Admin (List)
// @Tags [Admin] Admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param filter[name] query string false "Filter"
// @Param filter[email] query string false "Filter"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/admins [get]
func (controller *AdminControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	result, err := controller.adminService.GetAdmins(filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

// View godoc
// @Summary Admin (View)
// @Description Admin (View)
// @Tags [Admin] Admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/admins/{id} [get]
func (controller *AdminControllerImpl) View(c echo.Context) error {
	adminId, _ := strconv.Atoi(c.Param("id"))

	result, err := controller.adminService.GetAdminByID(adminId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
