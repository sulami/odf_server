package game

import "testing"

import "strings"

func TestParser(t *testing.T) {
	cases := []struct {
		in, want string
		fin bool
	}{
		{"START",	"OK WELCOME",	false},
		{"START stuff",	"OK WELCOME",	false},
		{"EXIT",	"OK BYE",	true},
		{"EXIT stuff",	"OK BYE",	true},
		{"not valid",	"ERR UNKWNCMD",	false},
	}

	g := DefaultGame{}

	for _, c := range cases {
		out, fin := g.Parse(strings.Fields(c.in))
		if c.want != out {
			t.Errorf("Parse(%q) == %q, want %q", c.in, out, c.want)
		}
		if c.fin != fin {
			t.Errorf("Parse fin (%q) == %t, want %t", c.in, fin,
				c.fin)
		}
	}
}

