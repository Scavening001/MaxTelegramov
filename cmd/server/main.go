package server

import (
	"MaxTelegramov/cmd/client"
)

type Server struct {
	clients  map[*client.Client]bool
	register chan *client.Client
}
