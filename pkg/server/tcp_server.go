package server

import (
	"fmt"
	"io"
	"net"
)

func StartTCPServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Failed to read data from connection:", err)
			return
		}

		if n == 0 {
			continue
		}

		fmt.Println("Received data:", string(buf[:n]))

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Failed to write data to connection:", err)
			return
		}
	}
}
