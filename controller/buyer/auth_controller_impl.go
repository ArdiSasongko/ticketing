package buyer_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// BuyerControllerImpl is the implementation of BuyerController interface
type BuyerControllerImpl struct {
	Service buyer_service.BuyerServiceInterface
}

// NewBuyerController creates a new instance of BuyerControllerImpl
func NewBuyerController(service buyer_service.BuyerServiceInterface) *BuyerControllerImpl {
	return &BuyerControllerImpl{Service: service}
}

// Register godoc
// @Summary Register a new buyer
// @Description Register a new buyer with the given details
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer body buyer_web.BuyerRequest true "Buyer details"
// @Success 201 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/auth/register [post]
func (controller *BuyerControllerImpl) Register(c echo.Context) error {
	newUser := new(buyer_web.BuyerRequest)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	result, err := controller.Service.Register(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Success", result))
}

// Login godoc
// @Summary Login a buyer
// @Description Login a buyer with email and password
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer body buyer_web.BuyerLoginRequest true "Login details"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/auth/login [post]
func (controller *BuyerControllerImpl) Login(c echo.Context) error {
	loginUser := new(buyer_web.BuyerLoginRequest)

	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(loginUser); err != nil {
		return err
	}

	userLogin, errLogin := controller.Service.Login(loginUser.Email, loginUser.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", userLogin))
}

// ViewMe godoc
// @Summary View buyer's information
// @Description View a buyer's information by ID
// @Tags buyer
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me [get]
func (controller *BuyerControllerImpl) ViewMe(c echo.Context) error {
	authId, _ := helper.GetAuthId(c)

	buyer, getBuyerErr := controller.Service.ViewMe(authId)
	if getBuyerErr != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, getBuyerErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", buyer))
}

// Update godoc
// @Summary Update a buyer's information
// @Description Update a buyer's information by ID
// @Tags buyer
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body buyer_web.BuyerUpdateRequest true "Updated buyer details"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me/update [put]
func (controller *BuyerControllerImpl) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	userID, _ := strconv.Atoi(claims.ID)

	updateUser := new(buyer_web.BuyerUpdateRequest)

	if err := c.Bind(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(updateUser); err != nil {
		return err
	}

	result, err := controller.Service.Update(userID, *updateUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
