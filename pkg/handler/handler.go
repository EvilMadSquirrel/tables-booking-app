package handler

import (
	"github.com/gin-gonic/gin"
	"tables-booking-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")

	{
		api.GET("/places", h.getPlaces)
		api.POST("/reserve", h.reserveTable)

		bookings := api.Group(":user_id/bookings")
		{
			bookings.GET("/", h.getAllBookings)
			bookings.GET(":id", h.getBookingByID)
			bookings.PUT(":id", h.updateBooking)
			bookings.DELETE(":id", h.deleteBooking)
		}
	}
	return router
}
