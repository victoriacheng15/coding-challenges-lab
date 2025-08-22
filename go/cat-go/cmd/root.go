package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Flags struct {
	numberNonBlank bool
	numberLines    bool
	showEnds       bool
}

var catFlags Flags

var rootCmd = &cobra.Command{
	Use:   "cat-go [flag] [file...]",
	Short: "Concatenate files and print to standard output",
	Long: `Concatenate and print files to standard output. It can be used to display the content of one or more files. If no file is specified, it reads from standard input.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files specified, reading from stdin")
			return
		}

		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
				continue
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				// TODO: Add flag logic here later
				// For now, just print the line
				fmt.Println(line)
			}

			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
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
	rootCmd.Flags().BoolVarP(&catFlags.numberNonBlank, "number-nonblank", "b", false, "Number non-empty output lines")
	rootCmd.Flags().BoolVarP(&catFlags.numberLines, "number", "n", false, "Number all output lines")
	rootCmd.Flags().BoolVarP(&catFlags.showEnds, "show-ends", "E", false, "Display $ at end of each line")
}


