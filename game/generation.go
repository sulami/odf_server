package game

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateUniverse() *[]Sector {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	u := make([]Sector, 50 + r.Intn(50))
	for i := range u {
		generateSector(&u[i], len(u))
	}
	return &u
}

func generateSector(s *Sector, uniSize int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s.worlds = make([]world, 2 + r.Intn(10))
	s.name = sectorNames[r.Intn(len(sectorNames))]
	for i := range s.worlds {
		generateWorld(&s.worlds[i], s.name + " " + strconv.Itoa(i+1))
	}
	s.x = r.Intn(uniSize)
	s.y = r.Intn(uniSize)
	// TODO race
}

func generateWorld(w *world, name string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w.name = name
	w.population = 1e6 + r.Intn(1e9)
	w.techLevel = 1 + r.Intn(6)
	if w.techLevel < r.Intn(10) {
		w.shipyard = true
	} else {
		w.shipyard = false
	}
	// TODO goods
}

func generateCaptain() *captain {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := &captain{
		name: humanFirstNames[r.Intn(len(humanFirstNames))] + " " +
			humanLastNames[r.Intn(len(humanLastNames))],
		experience: 0,
		rank: 0,
	}
	return c
}

var humanFirstNames = [...]string {
	"Rolaf",
	"Marten",
	"Justim",
	"Tery",
	"Henroy",
	"Waltev",
	"Atrip",
	"Mara",
	"Jana",
	"Dora",
	"Caria",
	"Donne",
	"Chera",
	"Mela",
	"Tine",
	"Diana",
}

var humanLastNames = [...]string {
	"Warder",
	"Grivis",
	"Risell",
	"Andex",
	"Rosson",
	"Pera",
	"Cooper",
	"Risach",
	"Harre",
	"Campbenn",
	"Hilley",
}

var sectorNames = [...]string {
	"Pike",
	"Xindi",
	"Mimir",
	"Ceani",
	"Calais",
	"Enyo",
	"Bani",
	"Aule",
	"Herschel",
	"Felis",
	"Ceberi",
}

