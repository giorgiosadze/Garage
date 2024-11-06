package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// nc localhost 8080
func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()
	fmt.Println("Listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn) // handle the connection
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 8) // Buffer to read 8 bytes
		conn.Read(buf)
		if strings.Contains(string(buf), "quit") {
			return
		}
		conn.Write(buf)
		fmt.Printf("Received 8 bytes: %s\n", buf)
	}
}
