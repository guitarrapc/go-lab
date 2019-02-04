## Refer

> https://stackoverflow.com/questions/54475723/which-is-better-to-get-the-last-x-character-of-a-golang-string

## Result

`getLastRune()` is much better choice.

1. getLastRune() is more than 4-5 times faster.
1. optimization remove allocation from both, but normally converting a string to []rune and back generally requires allocation.

## Benchmark

run benchmark.

```shell
go test -bench . -benchmem
```

```
goos: windows
goarch: amd64
pkg: github.com/guitarrapc/go-lab/stringsLab/sliceString
BenchmarkGetLastRune-8          100000000               22.6 ns/op             0 B/op          0 allocs/op
BenchmarkGetLastRune2-8         20000000               110 ns/op               0 B/op          0 allocs/op
PASS
```

run benchmark without optimization.

```shell
go test -gcflags '-N -l' -bench . -benchmem
```

result.

```
goos: windows
goarch: amd64
pkg: github.com/guitarrapc/go-lab/stringsLab/sliceString
BenchmarkGetLastRune-8          50000000                30.2 ns/op             0 B/op          0 allocs/op
BenchmarkGetLastRune2-8         10000000               133 ns/op              16 B/op          1 allocs/op
PASS
```