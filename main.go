package main

import (
	"log"
	"ws_notifications_email/interfaces/middleware/server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	server.InitServer().RunServer()

}
