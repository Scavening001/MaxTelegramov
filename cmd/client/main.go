package client

import "net"

type Client struct {
	conn net.Conn
	name string
	send chan string
}
