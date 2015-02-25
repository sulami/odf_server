package game

import (
)

type Game interface {
	Parse([]string) (string, bool)
	GenerateUniverse()
}

type DefaultGame struct {
	players []*player
	universe []*Sector
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

