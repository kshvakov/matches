# matches

```
go test -v -bench=. -cpu=1,2,4,8
```


strings.Join vs fmt.Sprintf 

```
BenchmarkStringsJoin     5000000           241 ns/op          32 B/op          2 allocs/op
BenchmarkStringsJoin-2   5000000           245 ns/op          32 B/op          2 allocs/op
BenchmarkStringsJoin-4   5000000           234 ns/op          32 B/op          2 allocs/op
BenchmarkStringsJoin-8  10000000           242 ns/op          32 B/op          2 allocs/op
BenchmarkSprintf         2000000           754 ns/op          64 B/op          4 allocs/op
BenchmarkSprintf-2       2000000           780 ns/op          64 B/op          4 allocs/op
BenchmarkSprintf-4       2000000           751 ns/op          64 B/op          4 allocs/op
BenchmarkSprintf-8       2000000           738 ns/op          64 B/op          4 allocs/op

```