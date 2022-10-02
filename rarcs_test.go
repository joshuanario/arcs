package arcs

import (
	"testing"
)

func TestRArcsPaths(t *testing.T) {
	arcs := testArcs
	root := NewRArc(arcs, "joshuanario.com")
	start := Path{}
	output := root.Paths(start)
	if len(output) != 3 {
		t.Fatalf("expected 3 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestRArcsHashQuery(t *testing.T) {
	arcs := testArcs
	root := NewRArc(arcs, "joshuanario.com")
	stimulus := "E"
	output := root.HashQuery(stimulus)
	if output == nil {
		t.Fatalf("expected non-nil locator;")
	}
	if output.Locator != stimulus {
		t.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
	}
}

func BenchmarkRArcs(b *testing.B) {
	arcs := stimulusArcs()
	for n := 0; n < b.N; n++ {
		root := NewRArc(arcs, "joshuanario.com")
		if root == nil {
			b.Fatalf("expected non-nil root;")
		}
	}
}

func BenchmarkRArcsHashQuery(b *testing.B) {
	arcs := stimulusArcs()
	root := NewRArc(arcs, "joshuanario.com")
	stimulus := "E"
	for n := 0; n < b.N; n++ {
		output := root.HashQuery(stimulus)
		if output == nil {
			b.Fatalf("expected non-nil locator;")
		}
		if output.Locator != stimulus {
			b.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
		}
	}
}

func BenchmarkRArcsPaths(b *testing.B) {
	arcs := stimulusArcs()
	root := NewRArc(arcs, "joshuanario.com")
	start := Path{}
	for n := 0; n < b.N; n++ {
		output := root.Paths(start)
		if len(output) != 28 {
			b.Fatalf("expected 28 paths; outcome %d;\n%v\n", len(output), output)
		}
	}
}
