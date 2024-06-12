package main

import (
	"bufio"
	"fmt"
	"net"
)

var clients = make(map[net.Conn]string)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clients[conn] = conn.RemoteAddr().String()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			delete(clients, conn)
			return
		}
		broadcast(message, conn)
	}
}

func broadcast(message string, sender net.Conn) {
	for conn := range clients {
		if conn != sender {
			fmt.Fprint(conn, message)
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Server started on :8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
