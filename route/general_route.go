package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/controller/general"
	general_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/general"
	"github.com/ArdiSasongko/ticketing_app/repository/general"
	"github.com/ArdiSasongko/ticketing_app/service/general"
	"github.com/labstack/echo/v4"
)

// RegisterGeneralRoutes registers general routes
// @Summary Register General Routes
// @Description Register all the general related routes
// @Tags general

func RegisterGeneralRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	generalEventQueryBuilder := general_query_builder.NewEventQueryBuilder(db)
	generalEventRepo := general_repository.NewEventRepository(generalEventQueryBuilder, db)
	generalEventService := general_service.NewEventService(generalEventRepo)
	generalEventController := general_controller.NewEventController(generalEventService)

	g := e.Group(prefix)

	eventRoute := g.Group("/events")
	eventRoute.GET("", generalEventController.List)
	eventRoute.GET("/:id", generalEventController.View)
}
