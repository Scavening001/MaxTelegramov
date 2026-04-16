package server

import (
	"bufio"
	"log"
	"net"
	"sync"
)

type Server struct {
	Address  string
	Listener net.Listener
	clients  []net.Conn
	mu       sync.Mutex
}

func NewServer(adress string) *Server {
	return &Server{
		Address: adress,
		clients: make([]net.Conn, 0),
	}
}

func (s *Server) removeClient(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, c := range s.clients {
		if c == conn {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
}

func (s *Server) broadcast(msg []byte, sender net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, client := range s.clients {
		if client == sender {
			continue
		}
		_, err := client.Write(append(msg, '\n'))
		if err != nil {
			log.Printf("Ошибка отправки сообщения клиенту %v: %v", client.RemoteAddr(), err)
		}
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
			continue
		}
		log.Printf("Добро пожаловать, %s", conn.RemoteAddr().String())

		s.mu.Unlock()
		s.clients = append(s.clients, conn)
		s.mu.Unlock()

		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer func() {
		s.removeClient(conn)
		conn.Close()
		log.Printf("Клиент отключился: %s", conn.RemoteAddr().String())
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Bytes()
		s.broadcast(msg, conn)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка чтения от %s: %v", conn.RemoteAddr().String(), err)
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
