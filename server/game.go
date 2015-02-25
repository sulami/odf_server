package server

type Game struct {
	players []*player
	universe []*Sector
}

func (g Game) Parse(cmd []string) (response string, fin bool) {
	switch cmd[0] {
	case "START":
		response = "OK STARTING"
	case "EXIT":
		response = "OK BYE"
		fin = true
	default:
		response = "ERR UNKWNCMD"
	}
	return
}

