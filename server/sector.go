package server

type Sector struct {
	name string
	x, y int
	race race
	worlds []*world
}

func (s *Sector) Info() string {
	return ""
}

