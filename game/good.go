package game

type good struct {
	name string
	baseWorth int
	techLevel int
	rarity int
	size int
}

func (g *good)String() string {
	return g.name
}

var tradeGoods = []good{
	good{
		name: "Computer Chips",
		baseWorth: 20,
		techLevel: 2,
		rarity: 2,
		size: 1,
	},
	good{
		name: "Food",
		baseWorth: 5,
		techLevel: 1,
		rarity: 1,
		size: 10,
	},
}

