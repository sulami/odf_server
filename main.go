package main

import (
	"flag"
)

import "github.com/sulami/odf_server/server"

func main() {
	port := flag.Int("port", 1339, "Port to listen on")
	flag.Parse()

	ser := server.GameServer{*port, false}
	e := ser.Listen()
	if e != nil {
		server.Log("Error: " + e.Error())
		return
	}

	e = ser.StopListening()
	if e != nil {
		server.Log("Error: " + e.Error())
		return
	}
}

