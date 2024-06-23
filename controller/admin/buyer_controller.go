package admin_controller

import "github.com/labstack/echo/v4"

type BuyerController interface {
	List(c echo.Context) error
	View(c echo.Context) error
}
