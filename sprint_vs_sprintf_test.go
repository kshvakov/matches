package matches

import (
	"fmt"
	"testing"
)

func BenchmarkSprintInt(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprint(i)
	}
}

func BenchmarkSprintfInt(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprintf("%d", i)
	}
}

func BenchmarkSprintFloat(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprint(float32(i) + 0.1)
	}
}

func BenchmarkSprintfFloat(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprintf("%f", float32(i)+0.1)
	}
}
