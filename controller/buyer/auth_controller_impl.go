package buyer_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
	"net/http"
)

// AuthControllerImpl is the implementation of AuthController interface
type AuthControllerImpl struct {
	authService buyer_service.AuthService
}

// NewAuthController creates a new instance of AuthControllerImpl
func NewAuthController(authService buyer_service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authService: authService,
	}
}

// Register godoc
// @Summary Auth (Register)
// @Description Auth (Register)
// @Tags [Buyer] Auth
// @Accept json
// @Produce json
// @Param buyer body buyer_web.RegisterBuyerRequest true "Buyer details"
// @Success 201 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/auth/register [post]
func (controller *AuthControllerImpl) Register(c echo.Context) error {
	userInput := new(buyer_web.RegisterBuyerRequest)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}
	if err := c.Validate(userInput); err != nil {
		return err
	}

	result, err := controller.authService.Register(*userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Register Success", result))
}

// Login godoc
// @Summary Auth (Login)
// @Description Auth (Login)
// @Tags [Buyer] Auth
// @Accept json
// @Produce json
// @Param buyer body buyer_web.LoginBuyerRequest true "Login Buyer Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/auth/login [post]
func (controller *AuthControllerImpl) Login(c echo.Context) error {
	loginUser := new(buyer_web.LoginBuyerRequest)

	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(loginUser); err != nil {
		return err
	}

	userLogin, errLogin := controller.authService.Login(loginUser.Email, loginUser.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Login Success", userLogin))
}

// View godoc
// @Summary Me (View)
// @Description Me (View)
// @Tags [Buyer] Me
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me [get]
func (controller *AuthControllerImpl) View(c echo.Context) error {
	authId, _ := helper.GetAuthId(c)

	buyer, getBuyerErr := controller.authService.ViewMe(authId)
	if getBuyerErr != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, getBuyerErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", buyer))
}

// Update godoc
// @Summary Me (Update)
// @Description Me (Update)
// @Tags [Buyer] Me
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body buyer_web.UpdateBuyerRequest true "Updated buyer details"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me/update [put]
func (controller *AuthControllerImpl) Update(c echo.Context) error {
	buyerId, _ := helper.GetAuthId(c)

	request := new(buyer_web.UpdateBuyerRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}
	if err := c.Validate(request); err != nil {
		return err
	}

	result, err := controller.authService.Update(buyerId, *request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
