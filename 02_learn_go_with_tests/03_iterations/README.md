# Iterations

To do stuff repeatedly in Go, you'll need for. In Go there are no while, do, until keywords, you can only use `for`

## Benchmarking
The `testing.B` gives you access to the cryptically named b.N.

* To run the benchmarks do `go test -bench=.`

```bash
go test -bench=.

goos: linux
goarch: amd64
pkg: iterations
cpu: Intel(R) Core(TM) i3-9100 CPU @ 3.60GHz
BenchmarkRepeat-4       11371281               100.5 ns/op
PASS
ok      iterations      1.256s
```

What `100.5` ns/op means is our function takes on average 100.5 nanoseconds to run (on my computer). Which is pretty ok! To test this it ran it `11371281` times.

>[NOTE]
 by default Benchmarks are run sequentially.

