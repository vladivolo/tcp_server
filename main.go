package main

import (
	"fmt"
	"strings"
	"tcp_server/tcpserver"
)

var (
	client_cnt  int
	message_cnt int
)

func main() {
	server := tcp_server.New("localhost:9999")

	server.OnNewClient(func(c *tcp_server.Client) {
		client_cnt++
		print_stat()
	})

	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		message_cnt++
		c.Send(fmt.Sprintf("%d\n", len(strings.TrimRight(message, "\r\n"))))
		print_stat()
	})

	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		client_cnt--
		print_stat()
	})

	server.Listen()
}

func print_stat() {
	fmt.Printf("%d\t%d\r", client_cnt, message_cnt)
}
