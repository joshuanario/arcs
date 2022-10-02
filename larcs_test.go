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

func BenchmarkLArcs(b *testing.B) {
	arcs := stimulusArcs()
	for n := 0; n < b.N; n++ {
		root := NewLArc(arcs, "joshuanario.com")
		if root == nil {
			b.Fatalf("expected non-nil root;")
		}
	}
}

func BenchmarkLArcsIterQuery(b *testing.B) {
	arcs := stimulusArcs()
	root := NewLArc(arcs, "joshuanario.com")
	stimulus := "E"
	for n := 0; n < b.N; n++ {
		output, _ := root.IterQuery(stimulus)
		if output == nil {
			b.Fatalf("expected non-nil locator;")
		}
		if output.Locator != stimulus {
			b.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
		}
	}
}

func BenchmarkLArcsPaths(b *testing.B) {
	arcs := stimulusArcs()
	root := NewLArc(arcs, "joshuanario.com")
	start := Path{}
	for n := 0; n < b.N; n++ {
		output := root.Paths(start)
		if len(output) != 28 {
			b.Fatalf("expected 28 paths; outcome %d;\n%v\n", len(output), output)
		}
	}
}
