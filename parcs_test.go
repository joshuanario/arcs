package arcs

import (
	"testing"
)

func TestPArcsPaths(t *testing.T) {
	arcs := testArcs
	root := NewPArc(arcs, "joshuanario.com")
	start := Path{}
	output := root.Paths(start)
	if len(output) != 3 {
		t.Fatalf("expected 3 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestPArcsIterQuery(t *testing.T) {
	arcs := testArcs
	root := NewPArc(arcs, "joshuanario.com")
	stimulus := "E"
	output, _ := root.IterQuery(stimulus)
	if output == nil || output.Locator != stimulus {
		t.Fatalf("expected \"E\" locator;")
	}
}
