package cat

import "testing"

func testDir() string {
	return "../../test-data"
}

func TestReadValidWarc(t *testing.T) {
	config := &config{
		fileName: testDir() + "/wikipedia/labrador_retriever_0.warc.gz"}
	_, err := listRecords(config, config.fileName)
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
}
