package game

import (
	"math/rand"
	"time"
)

func GenerateUniverse() []sector {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	u := make([]sector, 50 + r.Intn(50))
	for i := range u {
		go generateSector(&u[i], len(u))
	}
	return u
}

func generateSector(s *sector, uniSize int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s.worlds = make([]world, 2 + r.Intn(10))
	for i := range s.worlds {
		go generateWorld(&s.worlds[i])
	}
	s.x = r.Intn(uniSize)
	s.y = r.Intn(uniSize)
	// TODO name, race
}

func generateWorld(w *world) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w.population = 1e6 + r.Intn(1e9)
	w.techLevel = 1 + r.Intn(6)
	if w.techLevel < r.Intn(10) {
		w.shipyard = true
	} else {
		w.shipyard = false
	}
	// TODO name, goods
}

