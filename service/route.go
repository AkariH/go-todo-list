package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRoute() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		log.Println(c.Request.Body)
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
