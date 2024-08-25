package main

import (
	"todo-list/config"
	"todo-list/internal"
	"todo-list/service"
)

func main() {

	config.InitConfig()
	internal.InitDB()

	service.InitRoute()

}
