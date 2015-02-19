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

type Player struct {
	Experience int
	Level int

	// Base stats
	Health, MaxHealth int
	Stamina, MaxStamina int
	Mana, MaxMana int

	// Base attributes
	Strength int
	Agility int
	Endurance int
	Intelligence int
	Wisdom int
	Perception int
	Luck int
}

func RollAttribute() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20)
}

func NewPlayer() *Player {
	p := &Player{
		Experience: 0,
		Level: 1,
		Strength: RollAttribute(),
		Agility: RollAttribute(),
		Endurance: RollAttribute(),
		Intelligence: RollAttribute(),
		Wisdom: RollAttribute(),
		Perception: RollAttribute(),
		Luck: RollAttribute(),
	}

	p.MaxHealth = p.Endurance * p.Level / 2
	p.Health = p.MaxHealth
	p.MaxStamina = p.Strength * p.Level / 2
	p.Stamina = p.MaxStamina
	p.MaxMana = p.Wisdom * p.Level / 2
	p.Mana = p.MaxMana

	return p
}
