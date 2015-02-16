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
	n, err := c.Read(buf)
	if err != nil {
		log.Log("Error: " + err.Error())
	} else {
		log.Log("Incoming connection from " + c.RemoteAddr().String())
	}
	cmd := strings.Split(string(buf[:n]), " ")
	switch cmd[0] {
	case "LOGIN":
		if len(cmd) != 3 {
			c.Write([]byte("ERR ARGS\n"))
			c.Close()
		}
		if cmd[1] == "sulami" && cmd[2] == "" {
			c.Write([]byte("OK WELCOME\n"))
		} else {
			c.Write([]byte("ERR AUTH\n"))
		}
		// TODO find the user and try to auth him
	case "LOGOUT":
		c.Write([]byte("OK BYE\n"))
		c.Close()
	default:
		c.Write([]byte("ERR UNKWNCMD\n"))
		c.Close()
	}
}

