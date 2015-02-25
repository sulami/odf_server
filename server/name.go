package server

type nameSlice []string

func (slc *nameSlice) popItem(item int) string {
	s := *slc
	if item >= len(s) {
		return ""
	}
	r := s[item]
	s = append(s[:item], s[item+1:]...)
	*slc = s
	return r
}

var humanFirstNames = nameSlice{
	"Rolaf",
	"Marten",
	"Justim",
	"Tery",
	"Henroy",
	"Waltev",
	"Atrip",
	"Mara",
	"Jana",
	"Dora",
	"Caria",
	"Donne",
	"Chera",
	"Mela",
	"Tine",
	"Diana",
}

var humanLastNames = nameSlice{
	"Warder",
	"Grivis",
	"Risell",
	"Andex",
	"Rosson",
	"Pera",
	"Cooper",
	"Risach",
	"Harre",
	"Campbenn",
	"Hilley",
}

var sectorNames = nameSlice{
	"Pike",
	"Xindi",
	"Mimir",
	"Ceani",
	"Calais",
	"Enyo",
	"Bani",
	"Aule",
	"Herschel",
	"Felis",
	"Ceberi",
	"Guldek",
	"Wardan",
	"Geole",
	"Vidso",
	"Eosa",
}

