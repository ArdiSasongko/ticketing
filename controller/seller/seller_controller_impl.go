package seller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/model"
	web "github.com/ArdiSasongko/ticketing_app/model/web/seller"
	service "github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo"
)

type SellerControllerImpl struct {
	SellerService service.SellerService
}

func (controller *SellerControllerImpl) UpdateSeller(c echo.Context) error {

	seller := new(web.SellerUpdateServiceRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(seller); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	sellerUpdate, errSellerUpdate := controller.SellerService.UpdateSeller(*seller, id)

	if errSellerUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSellerUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "data berhasil diupdate", sellerUpdate))
}
