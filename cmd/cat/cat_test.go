package cat

import "testing"

func testDir() string {
	return "../../test-data"
}

func BenchmarkValidWarc(b *testing.B) {
	config := &config{
		fileName: testDir() + "/wikipedia/labrador_retriever_0.warc.gz"}
	_, err := listRecords(config, config.fileName)
	if err != nil {
		b.Errorf("Expected no error, got '%s'", err)
	}
}
