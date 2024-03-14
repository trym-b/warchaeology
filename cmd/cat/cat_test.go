package cat

import (
	"testing"
	"time"
)

func BenchmarkDummy(b *testing.B) {
	// This is a dummy test, it should be replaced with something more
	// meaningful in a later commit
	for i := 0; i < b.N; i++ {
		myFunc()
	}
}

func myFunc() {
	// This is a dummy function, it should be replaced with something more
	// meaningful in a later commit
	time.Sleep(1 * time.Second)
}
