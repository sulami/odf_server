package server

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Server interface {
	Listen() error
	StopListening() error
}

type GameServer struct {
	Port int
	Online bool
	Game Game
}

func (s *GameServer) Listen() (err error) {
	Log("Server starting up...")

	if s.Online {
		err = errors.New("Server is already online")
		return
	}

	Log("Generating universe...")
	s.Game = &DefaultGame{}
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
		go handleConnection(conn)
	}

	return
}

func (s *GameServer) StopListening() (err error) {
	Log("Stopping server...")

	if !s.Online {
		err = errors.New("Server is not running")
		return
	}

	s.Online = false
	Log("Server stopped")

	return
}

func handleConnection(conn net.Conn) {
	Log("Incoming connection from " + conn.RemoteAddr().String())
	NewClient(conn)
}

func Log(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04"), msg)
}

