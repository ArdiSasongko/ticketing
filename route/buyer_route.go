package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	buyer_controller "github.com/ArdiSasongko/ticketing_app/controller/buyer"
	order_controller "github.com/ArdiSasongko/ticketing_app/controller/order"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	buyer_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	buyer_repository "github.com/ArdiSasongko/ticketing_app/repository/buyer"
	history_repository "github.com/ArdiSasongko/ticketing_app/repository/history"
	order_repository "github.com/ArdiSasongko/ticketing_app/repository/order"
	buyer_service "github.com/ArdiSasongko/ticketing_app/service/buyer"
	order_service "github.com/ArdiSasongko/ticketing_app/service/order"
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
	orderRepository := order_repository.NewOrderRepository(db)
	orderService := order_service.NewOrderService(orderRepository, token)
	orderController := order_controller.NewOrderController(orderService)
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

	orderRoute := g.Group("/orders", middleware.JWTProtection())
	orderRoute.PUT("/orders/:id/pay", orderController.PayOrder)
}
