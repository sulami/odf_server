package server

import "errors"
import "strconv"
import "net"

import "github.com/sulami/odf_server/log"

type Server struct {
	Port	int
	Online	bool
}

func (s *Server) Listen() (err error) {
	log.Log("Server starting up...")

	if s.Online {
		log.Log("Error: Server is already online")
		err = errors.New("Server is already online")
		return
	}

	ln, e := net.Listen("tcp", ":" + string(s.Port))
	if e != nil {
		log.Log("Error: Could not bring up server")
		err = errors.New("Could not bring up server")
		return
	}

	for {
		conn, e:= ln.Accept()
		if e != nil {
			log.Log("Error: something")
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
		log.Log("Error: Server is not running")
		err = errors.New("Server is not running")
		return
	}

	s.Online = false
	log.Log("Server stopped")

	return
}

func handleConnection(c net.Conn) {
}

