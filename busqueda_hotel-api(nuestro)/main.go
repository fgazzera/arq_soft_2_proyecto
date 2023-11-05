package main

import (
	"busqueda_hotel-api/go_solr"
	"busqueda_hotel-api/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	ginRouter *gin.Engine
)

func main() {
	ginRouter = gin.Default()
	router.MapUrls(ginRouter)
	solrURL := "http://localhost:8983/solr"
	go_solr.InitSolrClient(solrURL)
	fmt.Println("Starting server")
	ginRouter.Run(":8090")
}
