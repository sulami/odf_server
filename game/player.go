package game

type player struct {
	name string
	money int
	homeworld *world
	worlds []*world
	fleets []*fleet
}

func initPlayer(name string, universe []Sector) player {
	p := player{
		name: name,
		money: 10000,
	}
	// TODO select a homeworld
	return p
}

