package router

import (
	hotelController "busqueda_hotel-api/controllers/hotel"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.POST("/hotel", hotelController.InsertHotel)
	//router.POST("/hotel_update/:id", hotelController.UpdateHotelById)

	fmt.Println("Finishing mappings configurations")
}
