package matches

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringsJoin(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		strings.Join([]string{"a", "[", "b", "]", "c"}, "")
	}
}

func BenchmarkSprintfJoin(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprintf("%s[%s]%s", "a", "b", "c")
	}
}
