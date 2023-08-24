package benchmark

import (
	"strconv"
	"testing"

	"github.com/flowstack/go-jsonparser"
)

func BenchmarkSetLarge(b *testing.B) {
	b.ReportAllocs()

	keyPath := make([]string, 20000)
	for i := range keyPath {
		keyPath[i] = "keyPath" + strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jsonparser.Set(largeFixture, largeFixture, keyPath...)
	}
}
