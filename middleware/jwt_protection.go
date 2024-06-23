package middleware

import (
	"net/http"
	"os"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "unauthorized", nil))
		},
	})
}

func auth(user enum.UserType, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(*jwt.Token).Claims.(*helper.JwtCustomClaims)
		if claims.Role != string(user) {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "Access forbidden"})
		}
		return next(c)
	}
}

func IsBuyer(next echo.HandlerFunc) echo.HandlerFunc {
	return auth(enum.UserTypeBuyer, next)
}

func IsSeller(next echo.HandlerFunc) echo.HandlerFunc {
	return auth(enum.UserTypeSeller, next)
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return auth(enum.UserTypeAdmin, next)
}
