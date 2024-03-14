package cat

import (
	"testing"
	"time"
)

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	return Fib(u-2) + Fib(u-1)
}
func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(10)
		time.Sleep(20 * time.Second)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}
func BenchmarkFib20WithAuxMetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
	b.ReportMetric(4.0, "auxMetricUnits")
}
