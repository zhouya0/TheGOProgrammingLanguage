package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

// 注意这三个channel都是无缓存的channel，所以都会阻塞
// 注意这个make(chan client)中的client也是一个 type client chan<- string
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// 可以看得出来这clients用map而不是slice，其实并不是clients中要存什么值，
// 而是clients便于增加，也可以方便删除。
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + "has arrived"
	// 包含channel的channel
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
