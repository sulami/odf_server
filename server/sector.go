package server

import "fmt"

type Sector struct {
	name string
	x, y int
	race race
	worlds []*world
}

// Return a string containing all of the sector information
// SECTOR <name> <x> <y> <race> <numWorlds> <[]worlds>
func (s *Sector) Info() string {
	retval := fmt.Sprintf(" SECTOR %s %d %d %s %d", s.name, s.x, s.y,
	                      s.race, len(s.worlds))

	for _, wor := range s.worlds {
		retval += wor.Info()
	}

	return retval
}

