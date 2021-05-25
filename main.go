package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.New()

	server.Use(gin.Logger(), gin.Recovery())

	server.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hola"}) })

	server.Run(":8080")
}