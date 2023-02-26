package main

import (
	"github.com/gin-gonic/gin",
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WOW, SUCCESS!!!!",
		})
	})
	r.Run()
}