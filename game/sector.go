package game

type Sector struct {
	name string
	x, y int
	race race
	worlds []*world
}

