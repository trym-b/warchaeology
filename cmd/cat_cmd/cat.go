package cat_cmd

import (
	"errors"

	"github.com/nlnwa/warchaeology/internal/filter"
	"github.com/nlnwa/warchaeology/internal/flag"
	cat_implementation "github.com/nlnwa/warchaeology/library/cat"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "cat",
		Short: "Concatenate and print warc files",
		Long:  ``,
		Example: `# Print all content from a WARC file
warc cat file1.warc.gz

# Pipe payload from record #4 into the image viewer feh
warc cat -n4 -P file1.warc.gz | feh -`,
		RunE: parseArgumentsAndCallCat,
	}

	cmd.Flags().Int64P(flag.Offset, "o", -1, flag.OffsetHelp)
	cmd.Flags().IntP(flag.RecordNum, "n", -1, flag.RecordNumHelp)
	cmd.Flags().IntP(flag.RecordCount, "c", 0, flag.RecordCountHelp+" Defaults to show all records except if -o or -n option is set, then default is one.")
	cmd.Flags().BoolP(flag.ShowWarcHeader, "w", false, flag.ShowWarcHeaderHelp)
	cmd.Flags().BoolP(flag.ShowProtocolHeader, "p", false, flag.ShowProtocolHeaderHelp)
	cmd.Flags().BoolP(flag.ShowPayload, "P", false, flag.ShowPayloadHelp)
	cmd.Flags().StringArray(flag.RecordId, []string{}, flag.RecordIdHelp)
	cmd.Flags().StringSliceP(flag.RecordType, "t", []string{}, flag.RecordTypeHelp)
	cmd.Flags().StringP(flag.ResponseCode, "S", "", flag.ResponseCodeHelp)
	cmd.Flags().StringSliceP(flag.MimeType, "m", []string{}, flag.MimeTypeHelp)

	return cmd
}

func parseArgumentsAndCallCat(cmd *cobra.Command, args []string) error {
	config := &cat_implementation.Config{}
	if len(args) == 0 {
		return errors.New("missing file name")
	}
	config.FileName = args[0]
	config.Offset = viper.GetInt64(flag.Offset)
	config.RecordCount = viper.GetInt(flag.RecordCount)
	config.RecordNum = viper.GetInt(flag.RecordNum)
	config.ShowWarcHeader = viper.GetBool(flag.ShowWarcHeader)
	config.ShowProtocolHeader = viper.GetBool(flag.ShowProtocolHeader)
	config.ShowPayload = viper.GetBool(flag.ShowPayload)

	if (config.Offset >= 0 || config.RecordNum >= 0) && config.RecordCount == 0 {
		config.RecordCount = 1
	}
	if config.Offset < 0 {
		config.Offset = 0
	}

	config.Filter = filter.NewFromViper()

	if !(config.ShowWarcHeader || config.ShowProtocolHeader || config.ShowPayload) {
		config.ShowWarcHeader = true
		config.ShowProtocolHeader = true
		config.ShowPayload = true
	}
	cat_implementation.ListFiles(config, config.FileName)
	return nil
}
