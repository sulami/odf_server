package server

type equipment interface {
	// for trading/equipment/etc
}

type hull struct {
	name string
	crewCapacity int
	weaponCapacity int
	shieldCapacity int
	cargoCapacity int
	cost int
}

type weapon struct {
	name string
	hullDamage int
	shieldDamage int
	powerDemand int
	size int
	cost int
}

type shield struct {
	name string
	strength int
	powerDemand int
	size int
	cost int
}

type generator struct {
	name string
	power int
	size int
	cost int
}

