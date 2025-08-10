package cmd

import (
	"fmt"
	"os"

	"wc-go/utils"

	"github.com/spf13/cobra"
)

type Flags struct {
	Lines bool
	Words bool
	Bytes bool
	Chars bool
}

var (
	flags Flags
)

var rootCmd = &cobra.Command{
	Use:   "wc-go [flags] [file...]",
	Short: "Count lines, words, and bytes in files or stdin",
	Long:  `wc-go is a simple command-line tool written in Go that counts lines, words, and bytes in input files or standard input. You can pass one or more files as arguments, or pipe input to it. If no flags are provided, wc-go prints all counts.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.ValidateArgs(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, filename := range args {
			data, err := os.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
				continue
			}

			switch {
			case flags.Bytes:
				fmt.Printf("%8d %s\n", utils.CountBytes(data), filename)
			case flags.Lines:
				fmt.Printf("%8d %s\n", utils.CountLines(data), filename)
			case flags.Words:
				fmt.Println("Printing words....")
			case flags.Chars:
				fmt.Println("Printing chars....")
			default:
				fmt.Println("Printing all....")
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
	rootCmd.Flags().BoolVarP(&flags.Lines, "lines", "l", false, "print the line counts")
	rootCmd.Flags().BoolVarP(&flags.Words, "words", "w", false, "print the word counts")
	rootCmd.Flags().BoolVarP(&flags.Bytes, "bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolVarP(&flags.Bytes, "chars", "m", false, "print the character counts")
}
