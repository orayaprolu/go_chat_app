package main

import "golang.org/x/net/websocket"

type Server struct {
	// Map of websocket connections
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server {
		conns: make(map[*websocket.Conn]bool)
	}
}

func (s *Server) handleWS() {

}

func main() {

}
