package arcs

import (
	"testing"
)

func TestCArcsPaths(t *testing.T) {
	arcs := testArcs
	root := NewCArc(arcs, "joshuanario.com")
	start := Path{}
	output := root.Paths(start)
	if len(output) != 3 {
		t.Fatalf("expected 3 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestCArcsIterQuery(t *testing.T) {
	arcs := testArcs
	root := NewCArc(arcs, "joshuanario.com")
	stimulus := "E"
	output, _ := root.IterQuery(stimulus)
	if output == nil {
		t.Fatalf("expected non-nil locator;")
	}
	if output.Locator != stimulus {
		t.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
	}
}

func BenchmarkCArcsIterQuery(b *testing.B) {
	arcs := stimulusArcs()
	for n := 0; n < b.N; n++ {
		root := NewCArc(arcs, "joshuanario.com")
		stimulus := "E"
		output, _ := root.IterQuery(stimulus)
		if output == nil {
			b.Fatalf("expected non-nil locator;")
		}
		if output.Locator != stimulus {
			b.Fatalf("expected \"E\" locator; outcome %s;\n%v\n", output.Locator, output)
		}
	}
}

func BenchmarkCArcsPaths(b *testing.B) {
	arcs := stimulusArcs()
	for n := 0; n < b.N; n++ {
		root := NewCArc(arcs, "joshuanario.com")
		start := Path{}
		output := root.Paths(start)
		if len(output) != 28 {
			b.Fatalf("expected 28 paths; outcome %d;\n%v\n", len(output), output)
		}
	}
}
