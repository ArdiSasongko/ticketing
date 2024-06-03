package seller

import (
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SellerControllerImpl struct {
	sellerService seller.SellerService
}

func NewSellerController(service seller.SellerService) *SellerControllerImpl {
	return &SellerControllerImpl{
		sellerService: service,
	}
}

func (controller *SellerControllerImpl) SaveSeller(c echo.Context) error {
	seller := new(seller.SellerServiceRequest)

	if err := c.Bind(seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	sellerUser, errSaveSeller := controller.sellerService.SaveSeller(*seller)

	if errSaveSeller != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveSeller.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "berhasil membuat user", sellerUser))
}

func (controller *SellerControllerImpl) LoginSeller(c echo.Context) error {
	seller := new(seller.SellerLoginRequest)

	if err := c.Bind(&seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	sellerRes, errLogin := controller.sellerService.LoginSeller(seller.Email, seller.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", sellerRes))
}
