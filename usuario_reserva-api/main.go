package main

import (
	"fmt"
	"usuario_reserva-api/router"
	"usuario_reserva-api/utils/db"

	"github.com/gin-gonic/gin"
)

var (
	ginRouter *gin.Engine
)

func main() {
	ginRouter = gin.Default()
	router.MapUrls(ginRouter)
	db.StartDbEngine()
	fmt.Println("Starting server")
	ginRouter.Run(":8090")
}
