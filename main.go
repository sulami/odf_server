package main

import "fmt"

import "github.com/sulami/odf_server/auth"

func main() {
	user := auth.User{"sulami", "robin"}
	fmt.Println(&user)
}

