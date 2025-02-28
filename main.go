package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/welcome/:nama", func(c *gin.Context) {
		name := c.Param("nama")
		c.String(http.StatusOK, "Selamat datang %s", name)
	})

	route.GET("/welcome/", func(c *gin.Context) {
		c.String(http.StatusOK, "Anonymous")
	})

	route.Run(":5000") // Listen and serve on 0.0.0.0:8080
}
