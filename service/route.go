package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo-list/internal"
)

func InitRoute() {
	r := gin.Default()
	if internal.DB == nil {
		log.Fatal("Database is still nil after initialization")
	} else {
		log.Println("Database initialized successfully")
	}

	r.GET("/ping", pingRoute)
	r.GET("/message/:id", getMessage)

	r.POST("/create", createMessage)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getMessage(c *gin.Context) {
	var DB = internal.DB
	id := c.Param("id")
	var message internal.Message
	log.Println(id)

	if err := DB.First(&message, id).Error; err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, message)
}

func createMessage(c *gin.Context) {
	var msg internal.Message
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println(err)
	}
	log.Println(msg)

}

func pingRoute(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "pong",
	})
}
