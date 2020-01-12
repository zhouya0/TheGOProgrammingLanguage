package main

import (
	"io"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		// It will be impossible to handle multiple connections
		// handleConn(conn)
		// This will make it be possible
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
