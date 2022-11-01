package main

import (
	"os"

	"github.com/ahmedabzk/go-microservices/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// This creates a router that can be used to define the build of the application.
	router := gin.Default()

	// load the templates
	router.LoadHTMLGlob("templates/*")

	// get the router
	routes.ArticlesRouter(router)

	router.Run(":" + port)
}
