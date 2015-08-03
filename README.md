# matches

```
go test -v -bench=. -cpu=1,2,4,8
```


strings.Join vs fmt.Sprintf vs concat (a + b + c)

```
BenchmarkStringsJoin     5000000           259 ns/op         160 B/op          2 allocs/op
BenchmarkStringsJoin-2   5000000           275 ns/op         160 B/op          2 allocs/op
BenchmarkStringsJoin-4   5000000           293 ns/op         160 B/op          2 allocs/op
BenchmarkStringsJoin-8   5000000           315 ns/op         160 B/op          2 allocs/op
BenchmarkSprintfJoin     2000000           722 ns/op         128 B/op          4 allocs/op
BenchmarkSprintfJoin-2   2000000           803 ns/op         128 B/op          4 allocs/op
BenchmarkSprintfJoin-4   2000000           824 ns/op         128 B/op          4 allocs/op
BenchmarkSprintfJoin-8   2000000           814 ns/op         128 B/op          4 allocs/op
BenchmarkConcat         10000000           165 ns/op          80 B/op          1 allocs/op
BenchmarkConcat-2       10000000           194 ns/op          80 B/op          1 allocs/op
BenchmarkConcat-4       10000000           178 ns/op          80 B/op          1 allocs/op
BenchmarkConcat-8       10000000           181 ns/op          80 B/op          1 allocs/op

```

fmt.Sprint vs fmt.Sprintf

```
BenchmarkSprintInt       5000000           262 ns/op          32 B/op          2 allocs/op
BenchmarkSprintInt-2     5000000           285 ns/op          32 B/op          2 allocs/op
BenchmarkSprintInt-4     5000000           278 ns/op          32 B/op          2 allocs/op
BenchmarkSprintInt-8     5000000           275 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt      5000000           301 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt-2    5000000           315 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt-4    5000000           309 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt-8    5000000           313 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat     3000000           455 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat-2   3000000           455 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat-4   3000000           475 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat-8   3000000           454 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat    3000000           553 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat-2  3000000           568 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat-4  3000000           566 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat-8  3000000           569 ns/op          32 B/op          2 allocs/op


```