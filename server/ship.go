package server

type ship struct {
	name string
	captain captain
	hull hull
	generator generator
	weapons []weapon
	shields []shield
	cargo map[good]int
}

func (s *ship) shieldValue() (v int) {
	for _, shield := range s.shields {
		v += shield.strength
	}
	return
}

func (s *ship) attackValue() (shield, hull int) {
	for _, weapon := range s.weapons {
		shield += weapon.shieldDamage
		hull += weapon.hullDamage
	}
	return
}

