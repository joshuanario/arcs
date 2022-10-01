package arcs

import (
	"testing"
)

var testArcs = []Arc{
	Arc{
		Arcrole: "joshuanario.com",
		Order:   3.0,
		From:    "A",
		To:      "B",
	},
	Arc{
		Arcrole: "resume.joshuanario.com",
		Order:   3.0,
		From:    "hireme",
		To:      "broadanddiverseskillset",
	},
	Arc{
		Arcrole: "resume.joshuanario.com",
		Order:   1.0,
		From:    "hireme",
		To:      "iamtalented",
	},
	Arc{
		Arcrole: "joshuanario.com",
		Order:   1.0,
		From:    "D",
		To:      "E",
	},
	Arc{
		Arcrole: "joshuanario.com",
		Order:   2.0,
		From:    "A",
		To:      "C",
	},
	Arc{
		Arcrole: "joshuanario.com",
		Order:   1.0,
		From:    "A",
		To:      "D",
	},
	Arc{
		Arcrole: "resume.joshuanario.com",
		Order:   2.0,
		From:    "hireme",
		To:      "longrecordofworkexperience",
	},
}

func stimulusArcs() []Arc {
	var ret []Arc
	ret = append(ret, testArcs...)
	var ch byte
	order := 0.0
	for ch = 'A'; ch <= 'Z'; ch++ {
		ret = append(ret, Arc{
			Arcrole: "joshuanario.com",
			Order:   order,
			From:    "root",
			To:      string([]byte{ch, ch}),
		})
		ret = append(ret, Arc{
			Arcrole: "resume.joshuanario.com",
			Order:   order,
			From:    "root",
			To:      string([]byte{ch, ch}),
		})
		order++
	}
	ret = append(ret, Arc{
		Arcrole: "joshuanario.com",
		Order:   order,
		From:    "ZZ",
		To:      "A",
	})
	return ret
}

func TestPaths(t *testing.T) {
	b := &controlArc{
		Locator: "B",
		Order:   1,
	}
	c := &controlArc{
		Locator: "C",
		Order:   1,
		Children: []*controlArc{
			b,
		},
	}
	stimulus := &controlArc{
		Locator: "A",
		Order:   1,
		Children: []*controlArc{
			b,
			c,
		},
	}
	start := Path{}
	output := Paths(stimulus, start)
	if len(output) != 2 {
		t.Fatalf("expected 2 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestManyPaths(t *testing.T) {
	f := &controlArc{
		Locator: "F",
		Order:   1,
	}
	d := &controlArc{
		Locator: "D",
		Order:   1,
	}
	b := &controlArc{
		Locator: "B",
		Order:   1,
		Children: []*controlArc{
			d,
		},
	}
	c := &controlArc{
		Locator: "C",
		Order:   2,
		Children: []*controlArc{
			b,
			f,
		},
	}
	a := &controlArc{
		Locator: "A",
		Order:   1,
		Children: []*controlArc{
			b,
			d,
		},
	}
	e := &controlArc{
		Locator: "F",
		Order:   3,
		Children: []*controlArc{
			d,
		},
	}
	stimulus := &controlArc{
		Locator: "root",
		Order:   1,
		Children: []*controlArc{
			a,
			c,
			e,
		},
	}
	start := Path{}
	output := Paths(stimulus, start)
	if len(output) != 5 {
		t.Fatalf("expected 5 paths; outcome %d;\n%v\n", len(output), output)
	}
}
