package server

import (
	"errors"
	"net"
	"strconv"
)

type Server struct {
	Port int
	Online bool
	Clients []*Client
	Game *Game
}

func (s *Server) Listen() (err error) {
	Log("Server starting up...")

	if s.Online {
		err = errors.New("Server is already online")
		return
	}

	Log("Generating universe...")
	s.Game = &Game{
		server: s,
		round: 0,
	}
	s.Game.GenerateUniverse()

	ln, err := net.Listen("tcp", ":" + strconv.Itoa(s.Port))
	if err != nil {
		return
	}

	s.Online = true
	Log("Listening on port " + strconv.Itoa(s.Port))

	for {
		conn, e:= ln.Accept()
		if e != nil {
			Log(e.Error())
		}
		go s.handleConnection(conn)
	}

	return
}

func (s *Server) StopListening() (err error) {
	Log("Stopping server...")

	if !s.Online {
		err = errors.New("Server is not running")
		return
	}

	s.Online = false
	Log("Server stopped")

	return
}

func (s *Server) handleConnection(conn net.Conn) {
	Log("Incoming connection from " + conn.RemoteAddr().String())
	c := NewClient(conn)
	c.game = s.Game
	s.Clients = append(s.Clients, c)
	c.Write("OK WELCOME")
	c.Write("QRY NAME")
}

func (s *Server) WriteAll(msg string) {
	for _, c := range s.Clients {
		c.Write(msg)
	}
}

