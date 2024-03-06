package internal

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMemoryAllocation(t *testing.T) {
	b := make([]byte, 2*1024*1024*1024)
	fmt.Println(len(b))
	pagesize := os.Getpagesize()
	for i := 0; i < 60; i++ {
		for j := 0; j < len(b); j += pagesize {
			b[j] = 42
		}
		time.Sleep(1 * time.Second)
	}
}
