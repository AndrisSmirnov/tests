package main

import (
	controllers "chat/app/controllers"
	nats_server "chat/app/servers/nats_server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	listeningPort := os.Getenv("SERVICE_PORT")

	// nats_server.Launch()
	nats_server.Publisher("chat_1", "admin", "user123", "ALOHA")

	router := controllers.InitControllers()
	fmt.Println("1_Chat")

	if err := router.Run(":" + listeningPort); err != nil {
		fmt.Println(err)
	}
	fmt.Println("2_Chat")
}
