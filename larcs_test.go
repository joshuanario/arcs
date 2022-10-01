package arcs

import (
	"testing"
)

func TestLArcsPaths(t *testing.T) {
	arcs := testArcs
	root := NewLArc(arcs, "joshuanario.com")
	start := Path{}
	output := root.Paths(start)
	if len(output) != 3 {
		t.Fatalf("expected 3 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestLArcsIterQuery(t *testing.T) {
	arcs := testArcs
	root := NewLArc(arcs, "joshuanario.com")
	stimulus := "E"
	output, _ := root.IterQuery(stimulus)
	if output == nil {
		t.Fatalf("expected non-nil locator;")
	}
	if output.Locator != stimulus {
		t.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
	}
}
