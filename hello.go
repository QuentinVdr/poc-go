package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"rsc.io/quote"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Optimization truth: %s", quote.Opt()),
		})
	})

	router.Run(":8080")
}
