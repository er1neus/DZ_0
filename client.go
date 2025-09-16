//go:build client

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		fmt.Println("Ошибка подключения:", error)
		os.Exit(1)
	}
	defer connection.Close()

	response, error := bufio.NewReader(connection).ReadString('\n')
	if error != nil {
		fmt.Println("Ошибка в считывании ответа:", error)
		os.Exit(1)
	}

	if response != "OK\n" {
		fmt.Println("Что-то пошло не так. Получен ответ:", response)
		os.Exit(1)
	}

	fmt.Println("Ответ сервера:", response)
}
