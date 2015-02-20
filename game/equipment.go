package game

type equipment interface {
	// for trading/equipment/etc
}

type hull struct {
	name string
	crewCapacity int
	weaponCapacity int
	shieldCapacity int
	cost int
}

type weapon struct {
	name string
	hull int
	shield int
	power int
	size int
	cost int
}

type shield struct {
	name string
	strength int
	power int
	size int
	cost int
}

type generator struct {
	name string
	power int
	size int
	cost int
}

