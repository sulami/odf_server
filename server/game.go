package server

type Universe []*Sector

type Game struct {
	server *Server
	players []*player
	universe *Universe
	round int
}

func (g *Game) Parse(cmd []string) (response string, fin bool) {
	switch cmd[0] {
	case "START":
		response = "OK STARTING"
		g.Start()
	case "NAME":
		response = "OK SET"
		// TODO set name accordingly
	case "EXIT":
		response = "OK BYE"
		fin = true
	default:
		response = "ERR UNKWNCMD"
	}
	return
}

func (g *Game) Start() {
	// Add one player per connected client
	g.server.WriteAll("UPD GAMESTART")
	for _, p := range g.server.Clients {
		player := initPlayer(p, "sulami", g.universe)
		g.players = append(g.players, player)
	}
}

func (g *Game) Round() {
	g.round += 1
	g.server.WriteAll("UPD UNIVERSE")
	// TODO package universe/game status into string
}

