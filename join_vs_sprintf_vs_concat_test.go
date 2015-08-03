package matches

import (
	"fmt"
	"strings"
	"testing"
)

var (
	_a = "aaaaaaaaaaaaaaaaaaaaaaaaaa"
	_b = "bbbbbbbbbbbbbbbbbbbbbbbbbb"
	_c = "cccccccccccccccccccccccccc"
)

func BenchmarkStringsJoin(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = strings.Join([]string{_a, "[", _b, "]", _c}, "")
	}
}

func BenchmarkSprintfJoin(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = fmt.Sprintf("%s[%s]%s", _a, _b, _c)
	}
}

func BenchmarkConcat(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = _a + "[" + _b + "]" + _c
	}
}
