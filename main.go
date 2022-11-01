package main

import (
	"net/http"

	"github.com/ahmedabzk/go-microservices/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	// This creates a router that can be used to define the build of the application.
	router := gin.Default()


	// load the templates
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			http.StatusOK,
			// use the index.html page
			"index.html",
			gin.H{
				"title":"Home Page",
			},
		)
	})

	routes.ArticlesRouter(router)
	
	router.Run()
}