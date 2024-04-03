package ls

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nlnwa/warchaeology/internal/filter"
	"github.com/spf13/afero"
)

func TestConfigReadFileWithError(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "test-data")
	warcWithErrors := filepath.Join(testDataDir, "samsung-with-error", "rec-33318048d933-20240317162652059-0.warc.gz")
	config := &conf{}
	config.filter = filter.NewFromViper()
	config.files = []string{warcWithErrors}
	_ = config.readFile(afero.NewOsFs(), warcWithErrors)
	// TODO: check that the result contains the expected values
}

func BenchmarkWriteLongStrings(b *testing.B) {
	// print extremely long strings instead
	// of the expected output.
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = writeLongStrings(os.Stdout)
	}

}

func writeLongStrings(out *os.File) error {
	_, err := out.WriteString("a")
	if err != nil {
		return err
	}
	for i := 0; i < 1000; i++ {
		_, err := out.WriteString("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		if err != nil {
			return err
		}
	}
	return nil
}
