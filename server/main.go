package main

import (
	"fmt"
	"net"
)

const PORT int = 6667

// const INSECURE_PORT int = 6697

func main() {
	fmt.Println("Welcome to the GoRC server :)")

	listener, err := net.Listen("tcp", ":6667")
	if err != nil {
		fmt.Println("Error listening on port 6667:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on port 6667")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}

		fmt.Println("Accepted connection from:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	b := make([]byte, 1024)
	i, err := conn.Read(b)
	if err != nil {
		fmt.Println("Unable to read from connection:", err)
	}

	mess := string(b[:i])
	fmt.Printf("Read %d bytes, Message Size: %d, Message: %s\n", i, len(mess), mess)
}
