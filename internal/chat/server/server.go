package server

import (
	"log"
	"net"
)

type Server struct {
	Address  string
	Listener net.Listener
}

func NewServer(adress string) *Server {
	return &Server{
		Address: adress,
	}
}

func (s *Server) Run() error {
	var err error
	s.Listener, err = net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}

	go s.acceptLoop()
	select {}
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Printf("Ошибка в подключении клиента: %v", err)
		}
		log.Printf("Добро пожаловать, %s", conn.RemoteAddr().String())
	}
}

// type Server struct {
// 	clients  map[*client.Client]bool
// 	register chan *client.Client
// 	unregister chan *client.Client
// 	broadcast chan []byte
// 	mu sync.RWMutex
// }

// func NewServer() *Server{
// 	return &Server{
// 		clients: make(map[*client.Client]bool),
// 		register: make(chan *client.Client),
// 		unregister: make(chan *client.Client),
// 		broadcast: make(chan []byte),
// 	}
// }

// func (s *Server) Run(){
// 	for{
// 		select{
// 		case cl := <-s.register:
// 			s.mu.Lock()
// 			s.clients[cl] = true
// 			s.mu.Unlock()
// 			s.broadcast <- []byte(cl.name + "Присоеденился") #
// 		case cl := <- s.unregister:
// 			s.mu.Lock()
// 			if _, ok := s.clients[cl];ok{
// 				delete(s.clients, cl)
// 				close(cl.Send)
// 			}

// 		}
// 	}
// }
