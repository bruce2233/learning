// generate_test.go
package benchmark

import (
	"math/rand"
	"testing"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

func benchmarkGenerate(scale int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(scale)
	}
}
func BenchmarkGenerate1000(b *testing.B) {
	benchmarkGenerate(1000, b)
}
func BenchmarkGenerate10000(b *testing.B) {
	benchmarkGenerate(10000, b)
}
func BenchmarkGenerate100000(b *testing.B) {
	benchmarkGenerate(100000, b)
}
func BenchmarkGenerate1000000(b *testing.B) {
	benchmarkGenerate(1000000, b)
}
