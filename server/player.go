package server

import (
	"math/rand"
	"time"
)

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
		homeworld: selectHomeworld(universe),
	}
	return p
}

func selectHomeworld(universe []Sector) *world {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	hs := universe[r.Intn(len(universe))]
	hw := hs.worlds[r.Intn(len(hs.worlds))]
	if hw.owner == nil {
		return hw
	} else {
		return selectHomeworld(universe)
	}
}

