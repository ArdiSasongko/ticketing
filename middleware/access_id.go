package middleware

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	seller_repository "github.com/ArdiSasongko/ticketing_app/repository/seller"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AccessUserID(repo seller_repository.EventRepositoryImpl) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
			}

			event, err := repo.GetEventByID(id)

			if err != nil {
				return c.JSON(http.StatusNotFound, helper.ResponseClient(http.StatusNotFound, err.Error(), nil))
			}

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*helper.JwtCustomClaims)
			userID, _ := strconv.Atoi(claims.ID)

			if event.SellerID != userID {
				return c.JSON(http.StatusForbidden, helper.ResponseClient(http.StatusForbidden, "You are not authorized to access this event", nil))
			}

			return next(c)
		}
	}
}
