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

func ReadFile(c *Config, fileName string) {
	wf, err := gowarc.NewWarcFileReader(fileName, c.Offset, gowarc.WithBufferTmpDir(viper.GetString(flag.TmpDir)))
	defer func() { _ = wf.Close() }()
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	num := 0
	count := 0

	for {
		wr, _, _, err := wf.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v, rec num: %v, Offset %v\n", err.Error(), strconv.Itoa(count), c.Offset)
			break
		}

		if !c.Filter.Accept(wr) {
			continue
		}

		// Find record number
		if c.RecordNum > 0 && num < c.RecordNum {
			num++
			continue
		}

		count++
		out := os.Stdout

		if c.ShowWarcHeader {
			// Write WARC record version
			_, err = fmt.Fprintf(out, "%v\r\n", wr.Version())
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			// Write WARC header
			_, err = wr.WarcHeader().Write(out)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			// Write separator
			_, err = out.WriteString("\r\n")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}

		if c.ShowProtocolHeader {
			if b, ok := wr.Block().(gowarc.ProtocolHeaderBlock); ok {
				_, err = out.Write(b.ProtocolHeaderBytes())
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			}
		}

		if c.ShowPayload {
			if pb, ok := wr.Block().(gowarc.PayloadBlock); ok {
				r, err := pb.PayloadBytes()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				_, err = io.Copy(out, r)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				r, err := wr.Block().RawBytes()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				_, err = io.Copy(out, r)
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

		if c.RecordCount > 0 && count >= c.RecordCount {
			break
		}
	}
	_, _ = fmt.Fprintln(os.Stderr, "Count: ", count)
}
