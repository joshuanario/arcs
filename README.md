# arcs

| arcs | sort | query |
|---|---|---|
| carcs | none (6400ns) | iter (inconsistent) |
| parcs | preemptive (6600ns) | iter (97ns) |
| larcs | lazy (707ns) | iter (97ns) |
| rarcs | preemptive (11000ns) | hash (18ns) |
| sarcs | lazy (690ns) | hash (12ns) |

carcs
```
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkCArcs-10    	  161680	      6451 ns/op	   17424 B/op	      66 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkCArcsIterQuery-10    	 9864424	       103.7 ns/op	       0 B/op	       0 allocs/opgoos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkCArcsPaths-10    	  380667	      2756 ns/op	    5256 B/op	      71 allocs/op
```

parcs
```
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkPArcs-10    	  179299	      6634 ns/op	   17560 B/op	      69 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkPArcsIterQuery-10    	10748806	        96.15 ns/op	       0 B/op	       0 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkPArcsPaths-10    	  438042	      2728 ns/op	    5256 B/op	      71 allocs/op
```

larcs
```
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkLArcs-10    	 1709046	       707.7 ns/op	    3536 B/op	       2 allocs/opgoos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkLArcsIterQuery-10    	12336084	        97.68 ns/op	       0 B/op	       0 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkLArcsPaths-10    	  418954	      2934 ns/op	    5256 B/op	      71 allocs/op
```

rarcs
```
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkRArcs-10    	   98157	     11407 ns/op	   25561 B/op	     114 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkRArcsHashQuery-10    	62999118	        18.86 ns/op	       0 B/op	       0 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkRArcsPaths-10    	  366218	      2740 ns/op	    5256 B/op	      71 allocs/op
```

sarcs
```
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkSArcs-10    	 1728267	       699.2 ns/op	    3552 B/op	       2 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkSArcsHashQuery-10    	100125984	        12.01 ns/op	       0 B/op	       0 allocs/op
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkSArcsPaths-10    	  347419	      2896 ns/op	    5256 B/op	      71 allocs/op
```