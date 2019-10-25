```
$ go test -v
=== RUN   TestSolver
--- PASS: TestSolver (0.00s)
=== RUN   TestSolver1
--- PASS: TestSolver1 (0.00s)
PASS
ok      interviewGo/lightOnSort 0.361s
```

```
$ go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: interviewGo/lightOnSort
BenchmarkSolver-4       10000000               157 ns/op             120 B/op          4 allocs/op
BenchmarkSolver2-4        500000              2032 ns/op             784 B/op         19 allocs/op
PASS
ok      interviewGo/lightOnSort 3.218s

```

