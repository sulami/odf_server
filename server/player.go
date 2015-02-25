package server

import (
	"math/rand"
	"time"
)

type player struct {
	client *Client
	name string
	money int
	homeworld *world
	worlds []*world
	fleets []*fleet
}

func initPlayer(client *Client, name string, universe *Universe) *player {
	p := &player{
		client: client,
		name: name,
		money: 10000,
		homeworld: selectHomeworld(universe),
	}
	return p
}

func selectHomeworld(universe *Universe) *world {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	hs := (*universe)[r.Intn(len(*universe))]
	hw := hs.worlds[r.Intn(len(hs.worlds))]
	if hw.owner == nil {
		return hw
	} else {
		return selectHomeworld(universe)
	}
}

