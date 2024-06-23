package seller_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SellerControllerImpl struct {
	sellerService seller_service.AuthService
}

func NewAuthController(service seller_service.AuthService) *SellerControllerImpl {
	return &SellerControllerImpl{
		sellerService: service,
	}
}

// Register godoc
// @Summary Auth (Register)
// @Description Auth (Register)
// @Tags [Seller] Auth
// @Accept json
// @Produce json
// @Param seller body seller_web.RegisterSellerRequest true "Register Seller Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/auth/register [post]
func (controller *SellerControllerImpl) Register(c echo.Context) error {
	seller := new(seller_web.RegisterSellerRequest)

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
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Register Success", sellerUser))
}

// Login godoc
// @Summary Auth (Login)
// @Description Auth (Login)
// @Tags [Seller] Auth
// @Accept json
// @Produce json
// @Param seller body seller_web.LoginSellerRequest true "Login Seller Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/auth/login [post]
func (controller *SellerControllerImpl) Login(c echo.Context) error {
	seller := new(seller_web.LoginSellerRequest)

	if err := c.Bind(&seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	sellerRes, errLogin := controller.sellerService.LoginSeller(seller.Email, seller.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Login Success", sellerRes))
}

// View godoc
// @Summary Me (View)
// @Description Me (View)
// @Tags [Seller] Me
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 404 {object} helper.ResponseClientModel
// @Router /seller/me [get]
func (controller *SellerControllerImpl) View(c echo.Context) error {
	id, _ := helper.GetAuthId(c)

	getSeller, errGetSeller := controller.sellerService.GetSeller(id)

	if errGetSeller != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseClient(http.StatusNotFound, errGetSeller.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", getSeller))
}

// Update godoc
// @Summary Me (Update)
// @Description Me (Update)
// @Tags [Seller] Me
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param seller body seller_web.UpdateSellerRequest true "Update Seller Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/me/update [put]
func (controller *SellerControllerImpl) Update(c echo.Context) error {
	userID, _ := helper.GetAuthId(c)
	seller := new(seller_web.UpdateSellerRequest)

	if err := c.Bind(seller); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}
	if err := c.Validate(seller); err != nil {
		return err
	}

	sellerUpdate, errSellerUpdate := controller.sellerService.UpdateSeller(*seller, userID)

	if errSellerUpdate != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errSellerUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Update Profile Success", sellerUpdate))
}
