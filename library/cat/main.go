package cat

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/nlnwa/gowarc"
	"github.com/nlnwa/warchaeology/internal/filter"
	"github.com/nlnwa/warchaeology/internal/flag"
	"github.com/spf13/viper"
)

type Config struct {
	Offset             int64
	RecordNum          int
	RecordCount        int
	FileName           string
	Filter             *filter.Filter
	ShowWarcHeader     bool
	ShowProtocolHeader bool
	ShowPayload        bool
}

func ListFiles(config *Config, fileName string) {
	warcFileReader, err := gowarc.NewWarcFileReader(fileName, config.Offset, gowarc.WithBufferTmpDir(viper.GetString(flag.TmpDir)))
	defer func() { _ = warcFileReader.Close() }()
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	num := 0
	count := 0

	for {
		warcRecord, _, _, err := warcFileReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v, rec num: %v, Offset %v\n", err.Error(), strconv.Itoa(count), config.Offset)
			break
		}

		if !config.Filter.Accept(warcRecord) {
			continue
		}

		// Find record number
		if config.RecordNum > 0 && num < config.RecordNum {
			num++
			continue
		}

		count++
		out := os.Stdout

		if config.ShowWarcHeader {
			// Write WARC record version
			_, err = fmt.Fprintf(out, "%v\r\n", warcRecord.Version())
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			// Write WARC header
			_, err = warcRecord.WarcHeader().Write(out)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			// Write separator
			_, err = out.WriteString("\r\n")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}

		if config.ShowProtocolHeader {
			if headerBlock, ok := warcRecord.Block().(gowarc.ProtocolHeaderBlock); ok {
				_, err = out.Write(headerBlock.ProtocolHeaderBytes())
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			}
		}

		if config.ShowPayload {
			if payloadBlock, ok := warcRecord.Block().(gowarc.PayloadBlock); ok {
				reader, err := payloadBlock.PayloadBytes()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				_, err = io.Copy(out, reader)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				reader, err := warcRecord.Block().RawBytes()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				_, err = io.Copy(out, reader)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			}
		}

		// Write end of record separator
		_, err = out.WriteString("\r\n\r\n")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if config.RecordCount > 0 && count >= config.RecordCount {
			break
		}
	}
	_, _ = fmt.Fprintln(os.Stderr, "Count: ", count)
}
