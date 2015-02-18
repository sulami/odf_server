package server

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

import "github.com/sulami/odf_server/client"

type Server interface {
	Listen() error
	StopListening() error
}

type GameServer struct {
	Port int
	Online bool
}

func (s *GameServer) Listen() (err error) {
	Log("Server starting up...")

	if s.Online {
		err = errors.New("Server is already online")
		return
	}

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
	client.NewClient(conn)
}

func Log(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04"), msg)
}

