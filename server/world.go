package server

import "fmt"

type world struct {
	name string
	population int
	techLevel int
	shipyard bool
	goods []*good
	owner *player
}

// Return a string containing all infos about this world
// WORLD <population> <techLevel> <shipyard> <owner> <[]goods>
func (w *world) Info() string {
	retval := fmt.Sprintf(" WORLD %d %d %t", w.population, w.techLevel,
	                      w.shipyard)

	if w.owner != nil {
		retval += w.owner.name
	}

	// TODO goods

	return retval
}

