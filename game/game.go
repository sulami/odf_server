package game

import (
	"math/rand"
	"time"
)

type Game interface {
	Parse([]string) (string, bool)
}

type DefaultGame struct {
}

func (g DefaultGame) Parse(cmd []string) (response string, fin bool) {
	switch cmd[0] {
	case "START":
		response = "OK WELCOME"
		// TODO start game routine
	case "EXIT":
		response = "OK BYE"
		fin = true
	default:
		response = "ERR UNKWNCMD"
	}
	return
}

/*
 * Down below: Game mechanic specific stuff - beware
 */

 type Entity interface {
	Attack(Entity, func(Entity, Entity) (bool, int))
	LoseHealth(int)

	Strength() int
	Agility() int
	Intelligence() int
	Perception() int
	Luck() int
 }

 //
 // PLAYER
 //
type Player struct {
	experience int
	level int

	// Base stats
	health, maxHealth int
	stamina, maxStamina int
	mana, maxMana int

	// Base attributes
	strength int
	agility int
	endurance int
	intelligence int
	wisdom int
	perception int
	luck int
}

func (p *Player) Attack(target Entity, atk func(Entity, Entity) (bool, int)) {
	hit, dmg := atk(p, target)
	if hit {
		target.LoseHealth(dmg)
	}
}

func (p *Player) LoseHealth(dmg int) {
	p.health = p.health - dmg
	if p.health <= 0 {
		// die
	}
}

//
// Player - Stat return methods
//
func (p *Player) Strength() int {
	return p.strength
}

func (p *Player) Agility() int {
	return p.agility
}

func (p *Player) Intelligence() int {
	return p.intelligence
}

func (p *Player) Perception() int {
	return p.perception
}

func (p *Player) Luck() int {
	return p.luck
}

func Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func RollAttribute() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20)
}

func NewPlayer() *Player {
	p := &Player{
		experience: 0,
		level: 1,
		strength: RollAttribute(),
		agility: RollAttribute(),
		endurance: RollAttribute(),
		intelligence: RollAttribute(),
		wisdom: RollAttribute(),
		perception: RollAttribute(),
		luck: RollAttribute(),
	}

	p.maxHealth = p.endurance * p.level / 2
	p.health = p.maxHealth
	p.maxStamina = p.strength * p.level / 2
	p.stamina = p.maxStamina
	p.maxMana = p.wisdom * p.level / 2
	p.mana = p.maxMana

	return p
}

func BasicMeleeAttack(source Entity, target Entity) (hit bool, dmg int) {
	if source.Strength() * Roll() >= target.Agility() * Roll() {
		hit = true
		dmg = source.Strength() * Roll() / 10
		if Roll() <= source.Luck() {
			dmg = dmg * 2
		}
	}
	return
}

//
// MONSTER
//
type Monster struct {
	worth int
	health, maxHealth int
	stamina, maxStamina int
	mana, maxMana int
	strength int
	agility int
	intelligence int
	perception int
	luck int
}

func NewMonster(lvl int) *Monster {
	m := &Monster{
		maxHealth: RollAttribute() * lvl / 2,
		maxStamina: RollAttribute() * lvl / 2,
		maxMana: RollAttribute() * lvl / 2,
		strength: RollAttribute(),
		agility: RollAttribute(),
		intelligence: RollAttribute(),
		perception: RollAttribute(),
		luck: RollAttribute(),
	}

	m.health = m.maxHealth
	m.stamina = m.maxStamina
	m.mana = m.maxMana
	m.worth = m.maxHealth + m.maxStamina + m.maxMana +
		(m.strength * lvl / 2) +
		(m.agility * lvl / 2) +
		(m.intelligence * lvl / 2) +
		(m.perception * lvl / 2)

	return m
}

//
// Monster - Stat return methods
//
func (m *Monster) Strength() int {
	return m.strength
}

func (m *Monster) Agility() int {
	return m.agility
}

func (m *Monster) Intelligence() int {
	return m.intelligence
}

func (m *Monster) Perception() int {
	return m.perception
}

func (m *Monster) Luck() int {
	return m.luck
}

