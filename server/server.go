package server

import (
	"fmt"
	"log"
	"net"
	"pidb/config"
)

type Server struct {
	address string
	port    int
	config  config.Config
}

func NewServer(address string, port int) *Server {
	return &Server{
		address: address,
		port:    port,
	}
}

func (server *Server) Run() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", server.address, server.port))
	if err != nil {
		log.Fatalln(err)
		return
	}
	socket, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for {
		tcpConn, err := socket.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
			return
		}
		client := NewClient(tcpConn)
		client.Run()
	}

}
