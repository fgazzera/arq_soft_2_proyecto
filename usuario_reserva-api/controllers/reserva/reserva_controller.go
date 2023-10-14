package reservaController

import (
	"net/http"
	"strconv"
	"usuario_reserva-api/dtos"
	service "usuario_reserva-api/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertReserva(c *gin.Context) {
	var reservaDto dtos.ReservaDto
	err := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservaDto, er := service.ReservaService.InsertReserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

func GetReservaById(c *gin.Context) {
	log.Debug("ID de reserva para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dtos.ReservaDto

	reservaDto, err := service.ReservaService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservaDto)
}

func GetReservasById(c *gin.Context) {
	log.Debug("ID de reserva para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservasDto dtos.ReservasDto

	reservasDto, err := service.ReservaService.GetReservasById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func GetDisponibilidad(c *gin.Context) {
	log.Debug("Disponibilidad de reservas para cargar: " + c.Param("id") + c.Param("AnioInicio") + c.Param("MesInicio") + c.Param("DiaInicio") + c.Param("AnioFinal") + c.Param("MesFinal") + c.Param("DiaFinal"))

	id, _ := strconv.Atoi(c.Param("id"))
	AnioInicio, _ := strconv.Atoi(c.Param("AnioInicio"))
	AnioFinal, _ := strconv.Atoi(c.Param("AnioFinal"))
	MesInicio, _ := strconv.Atoi(c.Param("MesInicio"))
	MesFinal, _ := strconv.Atoi(c.Param("MesFinal"))
	DiaInicio, _ := strconv.Atoi(c.Param("DiaInicio"))
	DiaFinal, _ := strconv.Atoi(c.Param("DiaFinal"))

	disponibilidad := service.ReservaService.GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal)

	c.JSON(http.StatusOK, disponibilidad)
}

func GetReservasByDate(c *gin.Context) {
	log.Debug("Reservas para cargar: " + c.Param("AnioInicio") + c.Param("MesInicio") + c.Param("DiaInicio") + c.Param("AnioFinal") + c.Param("MesFinal") + c.Param("DiaFinal"))

	AnioInicio, _ := strconv.Atoi(c.Param("AnioInicio"))
	AnioFinal, _ := strconv.Atoi(c.Param("AnioFinal"))
	MesInicio, _ := strconv.Atoi(c.Param("MesInicio"))
	MesFinal, _ := strconv.Atoi(c.Param("MesFinal"))
	DiaInicio, _ := strconv.Atoi(c.Param("DiaInicio"))
	DiaFinal, _ := strconv.Atoi(c.Param("DiaFinal"))

	var reservasDto dtos.ReservasDto

	reservasDto, err := service.ReservaService.GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}
