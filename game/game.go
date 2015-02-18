package game

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

