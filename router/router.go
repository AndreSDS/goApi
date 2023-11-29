package router

import "github.com/gin-gonic/gin" // usei o comando: go tidy para baixar as dependencias

func Initialize() { 
	// initialize the router
	router := gin.Default()
	//initialize the routes
	initializeRoutes(router)
	
	router.Run(":8080") // listen and serve on
}