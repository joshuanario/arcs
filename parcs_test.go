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
	if output == nil {
		t.Fatalf("expected non-nil locator;")
	}
	if output.Locator != stimulus {
		t.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
	}
}

func BenchmarkPArcs(b *testing.B) {
	arcs := stimulusArcs()
	for n := 0; n < b.N; n++ {
		root := NewPArc(arcs, "joshuanario.com")
		if root == nil {
			b.Fatalf("expected non-nil root;")
		}
	}
}

func BenchmarkPArcsIterQuery(b *testing.B) {
	arcs := stimulusArcs()
	root := NewPArc(arcs, "joshuanario.com")
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

func BenchmarkPArcsPaths(b *testing.B) {
	arcs := stimulusArcs()
	root := NewPArc(arcs, "joshuanario.com")
	start := Path{}
	for n := 0; n < b.N; n++ {
		output := root.Paths(start)
		if len(output) != 28 {
			b.Fatalf("expected 28 paths; outcome %d;\n%v\n", len(output), output)
		}
	}
}
