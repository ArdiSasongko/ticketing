package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	admin_web "github.com/ArdiSasongko/ticketing_app/model/web/admin"
	admin_service "github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	service admin_service.AdminServiceInterface
}

func NewAdminController(service admin_service.AdminServiceInterface) *AdminController {
	return &AdminController{
		service: service,
	}
}

func (controller *AdminController) Register(c echo.Context) error {
	newAdmin := new(admin_web.AdminRequest)

	if err := c.Bind(newAdmin); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newAdmin); err != nil {
		return err
	}

	result, err := controller.service.Register(*newAdmin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Register", result))
}

func (controller *AdminController) Login(c echo.Context) error {
	admin := new(admin_web.AdminLoginRequest)

	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(admin); err != nil {
		return err
	}

	result, err := controller.service.Login(admin.Email, admin.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Login", result))
}

func (controller *AdminController) GetBuyers(c echo.Context) error {
	result, err := controller.service.GetBuyers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Get Buyers", result))
}

func (controller *AdminController) GetSellers(c echo.Context) error {
	result, err := controller.service.GetSellers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Get Sellers", result))
}

func (controller *AdminController) GetBuyerByID(c echo.Context) error {
	buyerID, _ := strconv.Atoi(c.Param("id"))

	result, err := controller.service.GetBuyerByID(buyerID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Get Buyer By ID", result))
}

func (controller *AdminController) GetSellerByID(c echo.Context) error {
	sellerID, _ := strconv.Atoi(c.Param("id"))

	result, err := controller.service.GetSellerByID(sellerID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Get Seller By ID", result))
}
