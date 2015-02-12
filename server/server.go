package server

import "strconv"

import "github.com/sulami/odf_server/log"

type Server struct {
	Port	int
	Online	bool
}

func (s *Server) Listen() (err int) {
	log.Log("Server starting up...")

	if s.Online {
		log.Log("Error: Server is already online")
		err = 1
		return
	}

	log.Log("Listening on port " + strconv.Itoa(s.Port))

	return
}

