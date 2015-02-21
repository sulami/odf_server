package game

type player struct {
	name string
	money int
	homeworld world
	worlds []world
	fleets []fleet
}

