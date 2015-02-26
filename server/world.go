package server

type world struct {
	name string
	population int
	techLevel int
	shipyard bool
	goods []*good
	owner *player
}
