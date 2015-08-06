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

fmt.Sprint vs fmt.Sprintf vs strconv

```
BenchmarkSprintInt               5000000           266 ns/op          32 B/op          2 allocs/op
BenchmarkSprintInt-2             5000000           278 ns/op          32 B/op          2 allocs/op
BenchmarkSprintInt-4             5000000           276 ns/op          32 B/op          2 allocs/op
BenchmarkSprintInt-8             5000000           295 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt              5000000           319 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt-2            5000000           372 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt-4            5000000           330 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfInt-8            5000000           333 ns/op          32 B/op          2 allocs/op
BenchmarkStrconvItoa            20000000           125 ns/op          16 B/op          1 allocs/op
BenchmarkStrconvItoa-2          10000000           121 ns/op          16 B/op          1 allocs/op
BenchmarkStrconvItoa-4          20000000           130 ns/op          16 B/op          1 allocs/op
BenchmarkStrconvItoa-8          10000000           120 ns/op          16 B/op          1 allocs/op
BenchmarkSprintFloat             3000000           478 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat-2           3000000           550 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat-4           3000000           467 ns/op          32 B/op          2 allocs/op
BenchmarkSprintFloat-8           3000000           482 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat            2000000           779 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat-2          2000000           841 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat-4          2000000           752 ns/op          32 B/op          2 allocs/op
BenchmarkSprintfFloat-8          2000000           829 ns/op          32 B/op          2 allocs/op
BenchmarkStrconvFormatFloat      5000000           351 ns/op          48 B/op          2 allocs/op
BenchmarkStrconvFormatFloat-2    5000000           358 ns/op          48 B/op          2 allocs/op
BenchmarkStrconvFormatFloat-4    5000000           358 ns/op          48 B/op          2 allocs/op
BenchmarkStrconvFormatFloat-8    5000000           353 ns/op          48 B/op          2 allocs/op
```

gob vs encoding/json vs msgpack (gopkg.in/vmihailenco/msgpack.v2)
```
BenchmarkJsonMarshal          100000         14960 ns/op        3320 B/op          6 allocs/op
BenchmarkJsonMarshal-2        100000         18019 ns/op        3320 B/op          6 allocs/op
BenchmarkJsonMarshal-4        100000         17018 ns/op        3320 B/op          6 allocs/op
BenchmarkJsonMarshal-8        100000         16620 ns/op        3320 B/op          6 allocs/op
BenchmarkMsgpackMarshal       100000         21801 ns/op        1976 B/op          6 allocs/op
BenchmarkMsgpackMarshal-2      50000         23432 ns/op        1976 B/op          6 allocs/op
BenchmarkMsgpackMarshal-4     100000         25206 ns/op        1974 B/op          6 allocs/op
BenchmarkMsgpackMarshal-8     100000         23056 ns/op        1977 B/op          6 allocs/op
BenchmarkJsonUnmarshal         20000         65490 ns/op        4064 B/op        125 allocs/op
BenchmarkJsonUnmarshal-2       20000         69086 ns/op        4064 B/op        125 allocs/op
BenchmarkJsonUnmarshal-4       20000         67042 ns/op        4064 B/op        125 allocs/op
BenchmarkJsonUnmarshal-8       20000         68248 ns/op        4064 B/op        125 allocs/op
BenchmarkMsgpackUnmarshal      50000         32672 ns/op        2864 B/op         92 allocs/op
BenchmarkMsgpackUnmarshal-2    50000         34582 ns/op        2864 B/op         92 allocs/op
BenchmarkMsgpackUnmarshal-4    50000         33320 ns/op        2864 B/op         92 allocs/op
BenchmarkMsgpackUnmarshal-8    50000         33992 ns/op        2864 B/op         92 allocs/op
BenchmarkJsonEncoder          100000         13005 ns/op           8 B/op          1 allocs/op
BenchmarkJsonEncoder-2        100000         13014 ns/op           8 B/op          1 allocs/op
BenchmarkJsonEncoder-4        100000         12970 ns/op           8 B/op          1 allocs/op
BenchmarkJsonEncoder-8        100000         13130 ns/op           8 B/op          1 allocs/op
BenchmarkMsgpackEncoder       100000         18594 ns/op          64 B/op          2 allocs/op
BenchmarkMsgpackEncoder-2     100000         18669 ns/op          64 B/op          2 allocs/op
BenchmarkMsgpackEncoder-4     100000         18567 ns/op          64 B/op          2 allocs/op
BenchmarkMsgpackEncoder-8     100000         18540 ns/op          64 B/op          2 allocs/op
BenchmarkGobEncoder            50000         32983 ns/op        3744 B/op         54 allocs/op
BenchmarkGobEncoder-2          50000         36056 ns/op        3744 B/op         54 allocs/op
BenchmarkGobEncoder-4          50000         34600 ns/op        3744 B/op         54 allocs/op
BenchmarkGobEncoder-8          50000         37498 ns/op        3744 B/op         54 allocs/op
BenchmarkJsonDecoder           20000         69009 ns/op       10544 B/op        132 allocs/op
BenchmarkJsonDecoder-2         20000         74594 ns/op       10544 B/op        132 allocs/op
BenchmarkJsonDecoder-4         20000         73486 ns/op       10544 B/op        132 allocs/op
BenchmarkJsonDecoder-8         20000         76870 ns/op       10544 B/op        132 allocs/op
BenchmarkMsgpackDecoder        50000         32469 ns/op        7008 B/op         93 allocs/op
BenchmarkMsgpackDecoder-2      50000         37274 ns/op        7008 B/op         93 allocs/op
BenchmarkMsgpackDecoder-4      50000         37466 ns/op        7008 B/op         93 allocs/op
BenchmarkMsgpackDecoder-8      50000         35623 ns/op        7008 B/op         93 allocs/op
BenchmarkGobDecoder            10000        160588 ns/op       19740 B/op        420 allocs/op
BenchmarkGobDecoder-2          10000        118192 ns/op       19740 B/op        420 allocs/op
BenchmarkGobDecoder-4          10000        114654 ns/op       19741 B/op        420 allocs/op
BenchmarkGobDecoder-8          10000        113526 ns/op       19741 B/op        420 allocs/op

```