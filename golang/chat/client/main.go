package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: " + message)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, text)
	}
}
