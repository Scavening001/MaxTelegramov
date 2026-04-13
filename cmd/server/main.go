package main

import (
	"MaxTelegramov/internal/chat/server"
)

func main() {
	server := server.NewServer(":8080")
	server.Run()
}
