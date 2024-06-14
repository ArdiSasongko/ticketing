package seller_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	seller_web "github.com/ArdiSasongko/ticketing_app/model/web/seller"
	seller_service "github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type SellerControllerImpl struct {
	sellerService seller_service.SellerService
}

func NewSellerController(service seller_service.SellerService) *SellerControllerImpl {
	return &SellerControllerImpl{
		sellerService: service,
	}
}

func (controller *SellerControllerImpl) SaveSeller(c echo.Context) error {
	seller := new(seller_web.SellerServiceRequest)

	if err := c.Bind(seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(seller); err != nil {
		return err
	}

	sellerUser, errSaveSeller := controller.sellerService.SaveSeller(*seller)

	if errSaveSeller != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveSeller.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "berhasil membuat user", sellerUser))
}

func (controller *SellerControllerImpl) LoginSeller(c echo.Context) error {
	seller := new(seller_web.SellerLoginRequest)

	if err := c.Bind(&seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	sellerRes, errLogin := controller.sellerService.LoginSeller(seller.Email, seller.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", sellerRes))
}

func (controller *SellerControllerImpl) GetSeller(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getSeller, errGetSeller := controller.sellerService.GetSeller(id)

	if errGetSeller != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseClient(http.StatusNotFound, errGetSeller.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", getSeller))
}

func (controller *SellerControllerImpl) UpdateSeller(c echo.Context) error {

	seller := new(seller_web.SellerUpdateServiceRequest)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	userID, _ := strconv.Atoi(claims.ID)

	if err := c.Bind(seller); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	sellerUpdate, errSellerUpdate := controller.sellerService.UpdateSeller(*seller, userID)

	if errSellerUpdate != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errSellerUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "data berhasil diupdate", sellerUpdate))
}
