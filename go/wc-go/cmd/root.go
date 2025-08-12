package cmd

import (
	"fmt"
	"os"

	"wc-go/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wc-go [flags] [file...]",
	Short: "Count lines, words, and bytes in files or stdin",
	Long:  `wc-go is a simple command-line tool written in Go that counts lines, words, and bytes in input files or standard input. You can pass one or more files as arguments, or pipe input to it. If no flags are provided, wc-go prints all counts.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			data := utils.ReadContent("/dev/stdin", "stdin")
			utils.PrintCounts(data, "")
		} else {

			if err := utils.ValidateArgs(args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, filename := range args {
				data := utils.ReadContent(filename, fmt.Sprintf("file %s", filename))
				utils.PrintCounts(data, filename)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&utils.WcFlags.Lines, "lines", "l", false, "print the line counts")
	rootCmd.Flags().BoolVarP(&utils.WcFlags.Words, "words", "w", false, "print the word counts")
	rootCmd.Flags().BoolVarP(&utils.WcFlags.Bytes, "bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolVarP(&utils.WcFlags.Chars, "chars", "m", false, "print the character counts")
}
