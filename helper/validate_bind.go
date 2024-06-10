package helper

import (
	"fmt"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func BindAndValidate(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s field ini wajib diisi", err.Field())
				report.Code = http.StatusBadRequest
			case "email":
				report.Message = fmt.Sprintf("%s ini bukan email valid", err.Field())
				report.Code = http.StatusBadRequest
			}
		}
	}
	c.Logger().Error(report.Message)
	c.JSON(report.Code, model.ResponseToClient(report.Code, report.Message.(string), nil))
}
