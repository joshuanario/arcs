# arcs

| arcs | sort | query |
|---|---|---|
| carcs | none (6400ns) | iter (inconsistent) |
| parcs | preemptive (6600ns) | iter (97ns) |
| larcs | lazy (707ns) | iter (97ns) |
| rarcs | preemptive (11000ns) | hash (18ns) |
| sarcs | lazy (690ns) | hash (12ns) |


# benchmarks
```
goos: darwin
goarch: arm64
pkg: github.com/joshuanario/arcs
BenchmarkCArcs-10             	  161517	      6921 ns/op	   17424 B/op	      66 allocs/op
BenchmarkCArcsIterQuery-10    	10635060	       111.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkCArcsPaths-10        	  401482	      2901 ns/op	    5256 B/op	      71 allocs/op
BenchmarkLArcs-10             	 1681556	       707.3 ns/op	    3536 B/op	       2 allocs/op
BenchmarkLArcsIterQuery-10    	11220777	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLArcsPaths-10        	  432184	      2796 ns/op	    5256 B/op	      71 allocs/op
BenchmarkPArcs-10             	  169105	      6925 ns/op	   17560 B/op	      69 allocs/op
BenchmarkPArcsIterQuery-10    	11645016	       102.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkPArcsPaths-10        	  400548	      2887 ns/op	    5256 B/op	      71 allocs/op
BenchmarkRArcs-10             	  106341	     11352 ns/op	   25563 B/op	     114 allocs/op
BenchmarkRArcsHashQuery-10    	59329208	        19.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkRArcsPaths-10        	  406052	      2878 ns/op	    5256 B/op	      71 allocs/op
BenchmarkSArcs-10             	 1698196	       705.5 ns/op	    3552 B/op	       2 allocs/op
BenchmarkSArcsHashQuery-10    	87937592	        12.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkSArcsPaths-10        	  411038	      2902 ns/op	    5256 B/op	      71 allocs/op
```
