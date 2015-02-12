package main

import (
	"flag"
)

import "github.com/sulami/odf_server/server"

func main() {
	port := flag.Int("port", 1338, "Port to listen on")
	flag.Parse()

	server := server.Server{*port, false}
	e := server.Listen()
	if e != 0 {
		return
	}

	e = server.StopListening()
	if e != 0 {
		return
	}
}

