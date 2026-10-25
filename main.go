package main

import (
	"fmt"

	database "github.com/shaiklalahmed/notification-service/db"
	"github.com/shaiklalahmed/notification-service/service"
)

func main() {
	fmt.Println("First Go project")

	DB := database.NewService(2, 2)

	serv := service.NewServiceLayer(DB)
	value := serv.GetValueSer()

	fmt.Println(value)
	fmt.Println(value)
	fmt.Println(value)
	fmt.Println(value)

	fmt.Println(value)
	fmt.Println(value)
	fmt.Println(value)

}
