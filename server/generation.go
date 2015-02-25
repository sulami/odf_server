package server

import (
	"math/rand"
	"strconv"
	"time"
)

func (g *Game) GenerateUniverse() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	u := make([]*Sector, 5 + r.Intn(5))
	for i := range u {
		u[i] = generateSector(len(u))
	}
	g.universe = (*Universe)(&u)
}

func generateSector(uniSize int) *Sector {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := new(Sector)
	s.worlds = make([]*world, 2 + r.Intn(10))
	s.name = sectorNames.popItem(r.Intn(len(sectorNames)))
	for i := range s.worlds {
		s.worlds[i] = generateWorld(s.name + " " + strconv.Itoa(i+1))
	}
	s.x = r.Intn(uniSize)
	s.y = r.Intn(uniSize)
	s.race = raceList[r.Intn(len(raceList))]
	return s
}

func generateWorld(name string) *world {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w := new(world)
	w.name = name
	w.population = 1e6 + r.Intn(1e9)
	w.techLevel = 1 + r.Intn(6)
	if w.techLevel < r.Intn(10) {
		w.shipyard = true
	} else {
		w.shipyard = false
	}
	w.goods = make([]*good, 0)
	for i := 0; i < 1 + r.Intn(len(tradeGoods)); i++ {
		g := tradeGoods[r.Intn(len(tradeGoods))]
		inList := false
		// Duplication check
		for _, gl := range w.goods {
			if &g == gl {
				inList = true
			}
		}
		if !inList {
			// Tech Level check
			if g.techLevel <= w.techLevel {
				w.goods = append(w.goods, &g)
			}
		}
	}
	return w
}

func generateCaptain() *captain {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := &captain{
		name: humanFirstNames.popItem(r.Intn(len(humanFirstNames))) +
			" " +
			humanLastNames.popItem(r.Intn(len(humanLastNames))),
		experience: 0,
		rank: 0,
	}
	return c
}

