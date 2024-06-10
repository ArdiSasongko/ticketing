package buyer_controller

import "github.com/labstack/echo/v4"

type BuyerController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Update(c echo.Context) error
	GetAll(c echo.Context) error
	GetHistory(c echo.Context) error
}
