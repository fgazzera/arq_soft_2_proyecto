package hotel

import (
	"busqueda_hotel-api/dtos"
	service "busqueda_hotel-api/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHotelById(c *gin.Context) {
	id := c.Param("id")
	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}

func InsertHotel(c *gin.Context) {
	var hotelDto dtos.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Llama al servicio para agregar el hotel a Solr en lugar de la base de datos
	apiErr := service.HotelService.InsertHotel(hotelDto)

	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

/*func UpdateHotelById(c *gin.Context) {
	var hotelDto dtos.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	// Llama al servicio para actualizar el hotel en Solr en lugar de la base de datos
	updatedHotelDto, err := service.HotelService.UpdateHotelById(id, hotelDto)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, updatedHotelDto)
}*/
