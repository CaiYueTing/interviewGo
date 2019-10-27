#topic

- given n and return how much '1' in 11^n
- hint: pascal triangle 


```
$ go test -v
=== RUN   TestSolver
--- PASS: TestSolver (0.00s)
PASS
ok      interviewGo/11power     0.439s
```

```
go test -v -bench=. -benchmem
=== RUN   TestSolver
--- PASS: TestSolver (0.00s)
goos: windows
goarch: amd64
pkg: interviewGo/11power
BenchmarkSolver-4          50000             29779 ns/op           57200 B/op        483 allocs/op
PASS
ok      interviewGo/11power     2.227s

```