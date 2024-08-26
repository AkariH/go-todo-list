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
	r.POST("/update", updateMessage)

	r.DELETE("/delete", deleteMessage)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func deleteMessage(c *gin.Context) {

}

func updateMessage(c *gin.Context) {
	var (
		DB           = internal.DB
		message      internal.Message
		inputMessage internal.Message
	)
	c.BindJSON(&inputMessage)
	id := inputMessage.ID
	content := inputMessage.Content

	if err := DB.First(&message, id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, "record not found")
		return
	}
	log.Printf("before modified id = %v, content =%v", message.ID, message.Content)
	message.Content = content
	DB.Save(&message)
	log.Printf("after modified id = %v, content =%v", id, content)

	c.JSON(http.StatusOK, message)
}

func getMessage(c *gin.Context) {
	var DB = internal.DB
	id := c.Param("id")
	var message internal.Message
	log.Println(id)

	if err := DB.First(&message, id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, "record not found")
		return
	}
	c.JSON(http.StatusOK, message)
}

func createMessage(c *gin.Context) {
	var msg internal.Message
	var DB = internal.DB
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println(err)
	}

	DB.Create(&msg)
	c.JSON(http.StatusOK, &msg)

}

func pingRoute(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "pong",
	})
}
