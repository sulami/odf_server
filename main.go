package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

import "github.com/sulami/odf_server/auth"

func main() {
	port := flag.Int("port", 1338, "Port to listen on")
	flag.Parse()

	log("Using port: " + strconv.Itoa(*port))
	user := auth.User{"sulami", "robin"}
	log("User connected: " + user.Username)
}

func log(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04"), msg)
}

