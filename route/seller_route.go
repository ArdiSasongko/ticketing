package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/controller/seller"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	"github.com/ArdiSasongko/ticketing_app/query_builder/seller"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
)

// RegisterSellerRoutes mendaftarkan rute-rute untuk seller
// @Summary Register Seller Routes
// @Description Register all the seller related routes
// @Tags seller

func RegisterSellerRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	sellerAuthRepo := seller_repository.NewSellerRepository(db)
	sellerAuthService := seller_service.NewSellerService(sellerAuthRepo, token)
	sellerAuthController := seller_controller.NewSellerController(sellerAuthService)
	sellerEventQB := seller_query_builder.NewEventQueryBuilder(db)
	sellerEventRepo := seller_repository.NewEventRepository(sellerEventQB)
	sellerEventService := seller_service.NewEventService(sellerEventRepo)
	sellerEventController := seller_controller.NewEventController(sellerEventService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", sellerAuthController.SaveSeller)
	authRoute.POST("/login", sellerAuthController.LoginSeller)

	meRoute := g.Group("/me", middleware.JWTProtection())
	meRoute.GET("", sellerAuthController.GetSeller)
	meRoute.PUT("/update", sellerAuthController.UpdateSeller)

	eventRoute := g.Group("/events", middleware.JWTProtection())
	eventRoute.GET("", sellerEventController.GetEventList)
	eventRoute.GET("/:id", sellerEventController.ViewEvent, middleware.AccessUserID(*sellerEventRepo))
	eventRoute.POST("", sellerEventController.SaveEvents)
	eventRoute.PUT("/:id", sellerEventController.UpdateEvent, middleware.AccessUserID(*sellerEventRepo))
	eventRoute.PUT("/:id/update-status", sellerEventController.UpdateEventStatus, middleware.AccessUserID(*sellerEventRepo))
	eventRoute.PUT("/:event_id/tickets/:ticket_id/check-in", sellerEventController.CheckInTicket)
	eventRoute.DELETE("/:id", sellerEventController.DeleteEvent, middleware.AccessUserID(*sellerEventRepo))
}
