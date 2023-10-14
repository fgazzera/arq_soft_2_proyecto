package router

import (
	"fmt"
	reservaController "usuario_reserva-api/controllers/reserva"
	userController "usuario_reserva-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping

	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user/username/:username", userController.GetUserByUsername)
	router.GET("/user/email/:email", userController.GetUserByEmail)
	router.POST("/user", userController.InsertUser)

	router.POST("/reserva", reservaController.InsertReserva)
	router.GET("/reserva/:id", reservaController.GetReservaById)
	router.GET("/reservas/:id", reservaController.GetReservasById)
	router.GET("/disponibilidad/:id/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", reservaController.GetDisponibilidad)
	router.GET("/reservas-por-fecha/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", reservaController.GetReservasByDate)

	fmt.Println("Finishing mappings configurations")
}
