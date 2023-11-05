package main

import (
	"busqueda_hotel-api/cache"
	"busqueda_hotel-api/router"
	"busqueda_hotel-api/utils/db"
	"busqueda_hotel-api/utils/queue"
	"fmt"
	"sync" // Importa la librería sync
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ginRouter *gin.Engine
	wg        sync.WaitGroup // Declara un WaitGroup
)

func main() {
	time.Sleep(5 * time.Second)
	ginRouter = gin.Default()
	//Cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	ginRouter.Use(cors.New(config))

	router.MapUrls(ginRouter)

	cache.InitCache()

	err := db.InitDB()

	if err != nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return
	}

	fmt.Println("Starting server")

	// Incrementa el contador del WaitGroup antes de ejecutar StartReceiving
	wg.Add(1)
	go queue.StartReceiving()

	// Ejecuta el servidor HTTP en segundo plano
	go func() {
		defer wg.Done() // Decrementa el contador del WaitGroup cuando termine ginRouter.Run
		ginRouter.Run(":8001")
	}()

	// Espera a que ambas rutinas terminen
	wg.Wait()
}
