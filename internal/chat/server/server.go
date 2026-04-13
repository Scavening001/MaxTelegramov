package server

import "MaxTelegramov/internal/chat/client"

type Server struct {
	clients  map[*client.Client]bool
	register chan *client.Client
}
