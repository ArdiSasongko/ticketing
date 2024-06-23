package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/controller/buyer"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	"github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	"github.com/ArdiSasongko/ticketing_app/repository/buyer"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
)

// RegisterBuyerRoutes mendaftarkan rute-rute untuk buyer
// @Summary Register Buyer Routes
// @Description Register all the buyer related routes
// @Tags buyer

func RegisterBuyerRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	buyerAuthRepo := buyer_repository.NewAuthRepository(db)
	buyerAuthService := buyer_service.NewAuthService(buyerAuthRepo, token)
	buyerAuthController := buyer_controller.NewAuthController(buyerAuthService)
	buyerEventQB := buyer_query_builder.NewEventQueryBuilder(db)
	buyerEventRepo := buyer_repository.NewEventRepository(buyerEventQB, db)
	buyerEventService := buyer_service.NewEventService(buyerEventRepo)
	buyerEventController := buyer_controller.NewEventController(buyerEventService)
	buyerOrderQueryBuilder := buyer_query_builder.NewOrderQueryBuilder(db)
	buyerOrderRepo := buyer_repository.NewOrderRepository(buyerOrderQueryBuilder, db)
	buyerSellerRepo := buyer_repository.NewSellerRepository(db)
	buyerOrderService := buyer_service.NewOrderService(db, buyerOrderRepo, buyerSellerRepo)
	buyerOrderController := buyer_controller.NewOrderController(buyerOrderService)
	buyerTicketQueryBuilder := buyer_query_builder.NewTicketQueryBuilder(db)
	buyerTicketRepository := buyer_repository.NewTicketRepository(buyerTicketQueryBuilder, db)
	buyerTicketService := buyer_service.NewTicketService(buyerTicketRepository)
	buyerTicketController := buyer_controller.NewTicketController(buyerTicketService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", buyerAuthController.Register)
	authRoute.POST("/login", buyerAuthController.Login)

	g.Use(middleware.JWTProtection(), middleware.IsBuyer)

	meRoute := g.Group("/me")
	meRoute.GET("", buyerAuthController.View)
	meRoute.PUT("/update", buyerAuthController.Update)

	eventRoute := g.Group("/events")
	eventRoute.GET("", buyerEventController.List)
	eventRoute.GET("/:id", buyerEventController.View)

	orderRoute := g.Group("/orders")
	orderRoute.GET("", buyerOrderController.List)
	orderRoute.GET("/:id", buyerOrderController.View, middleware.AccessOrder(*buyerOrderRepo))
	orderRoute.POST("", buyerOrderController.Create)
	orderRoute.PATCH("/:id/pay", buyerOrderController.Pay, middleware.AccessOrder(*buyerOrderRepo))
	orderRoute.DELETE("/:id", buyerOrderController.Delete, middleware.AccessOrder(*buyerOrderRepo))

	ticketRoute := g.Group("/tickets")
	ticketRoute.GET("", buyerTicketController.List)
	ticketRoute.GET("/:id", buyerTicketController.View, middleware.AccessTicket(*buyerTicketRepository))
}
