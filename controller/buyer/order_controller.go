package buyer_controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	List(c echo.Context) error
	View(c echo.Context) error
	Create(c echo.Context) error
	Pay(c echo.Context) error
	Delete(c echo.Context) error
}
