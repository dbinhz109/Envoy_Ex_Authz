package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/public", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a public endpoint",
		})
	})

	router.GET("/private", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a private endpoint",
		})
	})

	router.Run(":3002")
}
