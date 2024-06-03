package seller

import "github.com/labstack/echo/v4"

type SellerController interface {
	SaveSeller(c echo.Context) error
	LoginSeller(c echo.Context) error
}
