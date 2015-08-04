package matches

import (
	"fmt"
	"strconv"
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

func BenchmarkStrconvItoa(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		strconv.Itoa(i)
	}
}

func BenchmarkSprintFloat(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprint(float64(i) + 0.1)
	}
}

func BenchmarkSprintfFloat(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		fmt.Sprintf("%f", float64(i)+0.1)
	}
}

func BenchmarkStrconvFormatFloat(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		strconv.FormatFloat(float64(i)+0.1, 'f', -1, 64)
	}
}
