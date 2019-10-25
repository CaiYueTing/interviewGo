#topic

- input an array of int means turing on the order of switch
- and the light is concatenation
```
input = [1,2,3,4,5] return 5
input = [1,2,3,5,4] return 4
input = [3,1,2,5,4] return 3
```

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

