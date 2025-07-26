package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("launching GoRC client")

	conn, err := net.Dial("tcp", "localhost:6667")
	if err != nil {
		fmt.Println("Error conecting to localhost:6667", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello from the client"))
	if err != nil {
		fmt.Println("error sending data:", err)
		return
	}

	fmt.Println("Successfully sent payload")
}
