package cat

import (
	"os"
	"testing"
)

func testDir() string {
	return "../../test-data"
}

func BenchmarkValidWarc(b *testing.B) {
	temp := os.Stdout
	temp2 := os.Stderr
	os.Stdout = nil
	os.Stderr = nil

	config := &config{
		fileName: testDir() + "/wikipedia/labrador_retriever_0.warc.gz"}
	_, err := listRecords(config, config.fileName)
	if err != nil {
		b.Errorf("Expected no error, got '%s'", err)
	}
	os.Stdout = temp
	os.Stderr = temp2
}
