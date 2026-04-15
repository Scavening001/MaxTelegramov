package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartClient() {
	_, err := connect("localhost:8080")
	if err != nil {
		fmt.Println("connection")
	}
}

func connect(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	//Прием собщений будет наверное GOрутиной :)
	//Отправка сообщений
	return conn, nil
}

func readUsername() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "anonymous"
	}

	return strings.TrimSpace(username)
}

type Client struct {
	conn net.Conn
	name string
	send chan string
}
