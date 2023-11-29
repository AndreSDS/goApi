package router

import "github.com/gin-gonic/gin" // usei o comando: go tidy para baixar as dependencias

func Initialize() { // função com letra maiuscula é exportada automaticamente
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensagem": "andré",
		})
	})
	router.Run(":8080") // listen and serve on
}