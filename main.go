package main

import (
	"flag"
	"fmt"
	"time"
)

import "github.com/sulami/odf_server/auth"

func main() {
	port := flag.Int("port", 1338, "Port to listen on")
	flag.Parse()

	fmt.Println(logtime(), "Using port:", *port)
	user := auth.User{"sulami", "robin"}
	fmt.Println(logtime(), "User connected:", user)
}

func logtime() string {
	return time.Now().Format("2006-01-02 15:05")
}

