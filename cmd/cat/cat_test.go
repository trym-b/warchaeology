package cat

import (
	"testing"
	"time"
)

func testDir() string {
	return "../../test-data"
}

//func BenchmarkValidWarc(b *testing.B) {
//	temp := os.Stdout
//	temp2 := os.Stderr
//	os.Stdout = nil
//	os.Stderr = nil
//
//	config := &config{
//		fileName: testDir() + "/wikipedia/labrador_retriever_0.warc.gz"}
//	_, err := listRecords(config, config.fileName)
//	if err != nil {
//		b.Errorf("Expected no error, got '%s'", err)
//	}
//	os.Stdout = temp
//	os.Stderr = temp2
//}

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	return Fib(u-2) + Fib(u-1)
}

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(10)
	}
	time.Sleep(1 * time.Second)
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
