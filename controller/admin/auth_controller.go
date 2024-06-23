package admin_controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	View(c echo.Context) error
	Update(c echo.Context) error
}
