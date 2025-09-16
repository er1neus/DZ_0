//go:build server

package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("OK\n"))
}

func main() {
	listener, error := net.Listen("tcp", ":8080")
	if error != nil {
		fmt.Println("Ошибка:", error)
		return
	}
	defer listener.Close()

	for {
		connection, error := listener.Accept()
		if error != nil {
			fmt.Println("Ошибка подключения к клиенту:", error)
			continue
		}
		go handleConnection(connection)
	}
}
