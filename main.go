package main

import (
	"flag"
	"strconv"
)

import "github.com/sulami/odf_server/auth"
import "github.com/sulami/odf_server/log"

func main() {
	port := flag.Int("port", 1338, "Port to listen on")
	flag.Parse()


	log.Log("Using port: " + strconv.Itoa(*port))
	user := auth.User{"sulami", "robin"}
	log.Log("User connected: " + user.Username)
}

