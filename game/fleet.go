package game

type fleet struct {
	name string
	commander captain
	location sector
	ships []ship
}

