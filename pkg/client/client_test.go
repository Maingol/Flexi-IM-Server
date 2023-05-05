package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"testing"
)

func TestConnectToTCPServer(t *testing.T) {
	serverAddr := "localhost:3000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		t.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	msg := "Hello, TCP server!"
	conn.Write([]byte(msg))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Received from server:", string(buf[:n]))
}

func TestConnectToWebSocketServer(t *testing.T) {
	serverAddr := "localhost:4000"

	conn, _, err := websocket.DefaultDialer.Dial("ws://"+serverAddr+"/ws", nil)
	if err != nil {
		t.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	msg := "Hello, ws server!"
	conn.WriteMessage(websocket.TextMessage, []byte(msg))

	_, resp, err := conn.ReadMessage()
	if err != nil {
		t.Fatal("Error reading message:", err)
	}
	fmt.Println("Received from server:", string(resp))
}
