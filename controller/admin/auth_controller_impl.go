package admin_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/web/admin"
	"github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
	"net/http"
)

// AuthControllerImpl is the implementation of AuthController interface
type AuthControllerImpl struct {
	authService admin_service.AuthService
}

// NewAuthController creates a new instance of AuthControllerImpl
func NewAuthController(authService admin_service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authService: authService,
	}
}

// Register godoc
// @Summary Auth (Register)
// @Description Auth (Register)
// @Tags [Admin] Auth
// @Accept json
// @Produce json
// @Param admin body admin_web.RegisterAdminRequest true "Register Admin Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/auth/register [post]
func (controller *AuthControllerImpl) Register(c echo.Context) error {
	newAdmin := new(admin_web.RegisterAdminRequest)

	if err := c.Bind(newAdmin); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newAdmin); err != nil {
		return err
	}

	result, err := controller.authService.Register(*newAdmin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Register Success", result))
}

// Login godoc
// @Summary Auth (Login)
// @Description Auth (Login)
// @Tags [Admin] Auth
// @Accept json
// @Produce json
// @Param admin body admin_web.LoginAdminRequest true "Login Admin Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/auth/login [post]
func (controller *AuthControllerImpl) Login(c echo.Context) error {
	admin := new(admin_web.LoginAdminRequest)

	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(admin); err != nil {
		return err
	}

	result, err := controller.authService.Login(*admin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Login Success", result))
}

// View godoc
// @Summary Me (View)
// @Description Me (View)
// @Tags [Admin] Me
// @Accept json
// @Produce json
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/me [get]
func (controller *AuthControllerImpl) View(c echo.Context) error {
	adminID, _ := helper.GetAuthId(c)

	result, err := controller.authService.ViewMe(adminID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

// Update godoc
// @Summary Me (Update)
// @Description Me (Update)
// @Tags [Admin] Me
// @Accept json
// @Produce json
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /admin/me/update [put]
func (controller *AuthControllerImpl) Update(c echo.Context) error {
	buyerId, _ := helper.GetAuthId(c)

	request := new(admin_web.UpdateAdminRequest)
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
