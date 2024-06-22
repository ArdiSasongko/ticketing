package admin_controller

import "github.com/labstack/echo/v4"

type AdminControllerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetBuyers(c echo.Context) error
	GetSellers(c echo.Context) error
	GetBuyerByID(c echo.Context) error
	GetSellerByID(c echo.Context) error
}
