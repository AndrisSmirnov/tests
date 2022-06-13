package main

import (
	"fmt"
	"os"
	controllers "user/app/controllers"
	nats_server "user/app/servers/nats_server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	go nats_server.Subscriber("chat_1")

	listeningPort := os.Getenv("SERVICE_PORT")

	router := controllers.InitControllers()
	fmt.Println("1_User")

	if err := router.Run(":" + listeningPort); err != nil {
		fmt.Println(err)
	}
	fmt.Println("2c_User")
}
