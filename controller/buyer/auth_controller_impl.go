package buyer_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	buyer_web "github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	buyer_service "github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

func (controller *BuyerControllerImpl) Update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

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

func (controller *BuyerControllerImpl) GetAll(c echo.Context) error {
	result, err := controller.Service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

func (controller *BuyerControllerImpl) GetHistory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	userID, _ := strconv.Atoi(claims.ID)

	result, err := controller.Service.GetHistory(userID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
