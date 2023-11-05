package router

import (
	hotelController "busqueda_hotel-api/controllers/hotel"
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Hotel Mapping
	router.GET("/hotel", hotelController.GetHotels)
	router.GET("/hotel/:id", hotelController.GetHotel)
	router.GET("/disponibilidad/:fechainicio/:fechafinal/:ciudad", hotelController.GetDisponibilidad)
	router.GET("/disponibilidad/:fechainicio/:fechafinal/", hotelController.GetDisponibilidad)
	router.POST("/hotel", hotelController.CreateHotel)

	fmt.Println("Finishing mappings configurations")
}
