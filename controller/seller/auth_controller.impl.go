package seller_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
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

// SaveSeller godoc
// @Summary Create a new seller
// @Description Create a new seller with the input payload
// @Tags seller
// @Accept json
// @Produce json
// @Param seller body seller_web.SellerServiceRequest true "Create Seller Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/auth/register [post]
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
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully created user", sellerUser))
}

// LoginSeller godoc
// @Summary Login a seller
// @Description Login with the input payload
// @Tags seller
// @Accept json
// @Produce json
// @Param seller body seller_web.SellerLoginRequest true "Login Seller Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/auth/login [post]
func (controller *SellerControllerImpl) LoginSeller(c echo.Context) error {
	seller := new(seller_web.SellerLoginRequest)

	if err := c.Bind(&seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	sellerRes, errLogin := controller.sellerService.LoginSeller(seller.Email, seller.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Login successful", sellerRes))
}

// GetSeller godoc
// @Summary Get seller information by ID
// @Description Get seller information based on seller ID
// @Tags seller
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 404 {object} helper.ResponseClientModel
// @Router /seller/me [get]
func (controller *SellerControllerImpl) GetSeller(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getSeller, errGetSeller := controller.sellerService.GetSeller(id)

	if errGetSeller != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseClient(http.StatusNotFound, errGetSeller.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", getSeller))
}

// UpdateSeller godoc
// @Summary Update a seller
// @Description Update a seller by its ID
// @Tags seller
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param seller body seller_web.SellerUpdateServiceRequest true "Update Seller Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/me/update [put]
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

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Data successfully updated", sellerUpdate))
}
