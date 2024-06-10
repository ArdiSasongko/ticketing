package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/controller/buyer"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	"github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/history"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
)

func RegisterBuyerRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	historyRepo := history_repository.NewHistoryRepoImpl(db)
	buyerAuthRepo := buyer_repository.NewBuyerRepository(db)
	buyerAuthService := buyer_service.NewBuyerService(buyerAuthRepo, token, historyRepo)
	buyerAuthController := buyer_controller.NewBuyerController(buyerAuthService)
	buyerEventQB := buyer_query_builder.NewEventQueryBuilder(db)
	buyerEventRepo := buyer_repository.NewEventRepository(buyerEventQB)
	buyerEventService := buyer_service.NewEventService(buyerEventRepo)
	buyerEventController := buyer_controller.NewEventController(buyerEventService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", buyerAuthController.Register)
	authRoute.POST("/login", buyerAuthController.Login)

	meRoute := g.Group("/me", middleware.JWTProtection())
	meRoute.POST("/update", buyerAuthController.Update)
	meRoute.GET("/buyers", buyerAuthController.GetAll)
	meRoute.GET("/history", buyerAuthController.GetHistory)

	eventRoute := g.Group("/events", middleware.JWTProtection())
	eventRoute.GET("", buyerEventController.GetEventList)
}
