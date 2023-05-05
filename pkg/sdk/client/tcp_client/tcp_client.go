package tcp_client

import (
	"net"
)

func ConnectToTCPServer(msg string) string {
	serverAddr := "localhost:3000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return ""
	}
	defer conn.Close()

	conn.Write([]byte(msg))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	resMsg := "Received from server:" + string(buf[:n])
	return resMsg
}
