package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"todo-list/internal"
)

var (
	DB *gorm.DB
)

func InitRoute() {
	r := gin.Default()
	DB = internal.GetDB()
	if DB == nil {
		log.Fatal("Database is still nil after initialization")
	} else {
		log.Println("Database initialized successfully")
	}

	r.GET("/ping", pingRoute)
	r.GET("/message/:id", getMessage)

	r.POST("/create", createMessage)
	r.POST("/update", updateMessage)

	r.DELETE("/delete/:id", deleteMessage)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func deleteMessage(c *gin.Context) {
	var message internal.Message
	id := c.Param("id")

	if err := DB.First(&message, id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, "record not found")
		return
	}
	c.JSON(http.StatusOK, message)
	DB.Delete(&message, id)

}

func updateMessage(c *gin.Context) {
	var (
		inputMessage internal.Message
		message      internal.Message
	)
	err := c.BindJSON(&inputMessage)
	if err != nil {
		log.Println(err)
	}
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
	id := c.Param("id")
	var message internal.Message
	log.Printf("id is %v\n", id)

	if err := DB.First(&message, id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, "record not found")
		return
	}
	c.JSON(http.StatusOK, message)
}

func createMessage(c *gin.Context) {

	var message internal.Message
	err := c.BindJSON(&message)
	if err != nil {
		log.Println(err)
	}

	DB.Create(&message)
	c.JSON(http.StatusOK, &message)

}

func pingRoute(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "pong",
	})
}
