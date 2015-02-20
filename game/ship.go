package game

type ship struct {
	name string
	captain captain
	hull hull
	generator generator
	weapons []weapon
	shields []shield
}

func (s *ship) shieldValue() (v int) {
	for _, shield := range s.shields {
		v += shield.strength
	}
	return
}
