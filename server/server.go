package server

import (
	"errors"
	"strconv"
	"strings"
	"net"
)

import "github.com/sulami/odf_server/log"

type Server struct {
	Port	int
	Online	bool
}

func (s *Server) Listen() (err error) {
	log.Log("Server starting up...")

	if s.Online {
		err = errors.New("Server is already online")
		return
	}

	ln, e := net.Listen("tcp", ":" + strconv.Itoa(s.Port))
	if e != nil {
		err = e
		return
	}

	for {
		conn, e:= ln.Accept()
		if e != nil {
			log.Log(e.Error())
		}
		go handleConnection(conn)
	}

	s.Online = true
	log.Log("Listening on port " + strconv.Itoa(s.Port))

	return
}

func (s *Server) StopListening() (err error) {
	log.Log("Stopping server...")

	if !s.Online {
		err = errors.New("Server is not running")
		return
	}

	s.Online = false
	log.Log("Server stopped")

	return
}

func handleConnection(c net.Conn) {
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if err != nil {
		log.Log("Error: " + err.Error())
	} else {
		log.Log("Incoming connection from " + c.RemoteAddr().String())
	}
	cmd := strings.Split(string(buf), " ")
	switch cmd[0] {
	case "LOGIN":
		if len(cmd) != 3 {
			c.Write([]byte("ERRARGS"))
			c.Close()
		}
		// TODO find the user and try to auth him
		break;
	case "LOGOUT":
		c.Close()
		break;
	default:
		c.Write([]byte("ERRUNKWNCMD"))
		c.Close()
	}
}

