package server

import (
	"fmt"
	"log"
	"net"
	"pidb/config"
)

type Client struct {
	conn       *net.TCPConn
	WriterPipe chan []byte
	config     config.Config
}

func NewClient(conn *net.TCPConn) *Client {
	return &Client{
		conn:       conn,
		WriterPipe: make(chan []byte),
	}
}

func (client *Client) Run() {
	go client.reader()
	go client.writer()
}

func (client *Client) reader() {
	buffer := make([]byte, client.config.ServerConf.ClientReadBufferSize)
	for {
		n, err := client.conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		read := buffer[:n]
		fmt.Println(read)
	}
}

func (client *Client) writer() {
	for {
		select {
		case message := <-client.WriterPipe:
			n, err := client.conn.Write(message)
			if err != nil {
				log.Println(err)
			}
			if n != len(message) {
				log.Println("Length does not match")
			}
		}
	}
}
