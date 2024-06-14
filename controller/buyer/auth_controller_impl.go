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

type BuyerControllerImpl struct {
	Service buyer_service.BuyerServiceInterface
}

func NewBuyerController(service buyer_service.BuyerServiceInterface) *BuyerControllerImpl {
	return &BuyerControllerImpl{Service: service}
}

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

func (controller *BuyerControllerImpl) ViewMe(c echo.Context) error {
	authId, _ := helper.GetAuthId(c)

	buyer, getBuyerErr := controller.Service.ViewMe(authId)
	if getBuyerErr != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, getBuyerErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", buyer))
}

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
