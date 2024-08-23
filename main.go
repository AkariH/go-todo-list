package main

import (
	"log"
	"todo-list/config"
	"todo-list/internal"
	"todo-list/service"
)

func main() {
	log.Println("hello world")

	config.InitConfig()
	internal.InitDB()

	service.InitRoute()

}
