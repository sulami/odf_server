package server

type fleet struct {
	name string
	commander captain
	location world
	ships []ship
}

