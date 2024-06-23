package middleware

import (
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
	"github.com/labstack/echo/v4"
)

func AccessEvent(repo seller_repository.EventRepositoryImpl) echo.MiddlewareFunc {
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

			userID, _ := helper.GetAuthId(c)
			if event.SellerID != userID {
				return c.JSON(http.StatusForbidden, helper.ResponseClient(http.StatusForbidden, "You are not authorized to access this event", nil))
			}

			return next(c)
		}
	}
}

func AccessOrder(repo buyer_repository.OrderRepositoryImpl) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
			}

			history, err := repo.GetHistory(id)

			if err != nil {
				return c.JSON(http.StatusNotFound, helper.ResponseClient(http.StatusNotFound, err.Error(), nil))
			}

			userID, _ := helper.GetAuthId(c)
			if history.BuyerIDFK != userID {
				return c.JSON(http.StatusForbidden, helper.ResponseClient(http.StatusForbidden, "You are not authorized to access this order", nil))
			}

			return next(c)
		}
	}
}
