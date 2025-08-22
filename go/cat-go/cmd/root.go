package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Flags struct {
	NumberNonBlank bool
	NumberLines    bool
	ShowEnds       bool
}

var catFlags Flags

// formatLine applies the formatting flags to a line and returns the formatted output
func formatLine(line string, lineNumber *int) string {
	originalLine := line

	// Add $ at end if showEnds flag is set
	if catFlags.ShowEnds {
		line += "$"
	}

	if catFlags.NumberNonBlank {
		if len(originalLine) > 0 { // Check original line for emptiness
			result := fmt.Sprintf("%6d\t%s", *lineNumber, line)
			*lineNumber++
			return result
		} else {
			return line // This will just be "$" if showEnds is true
		}
	} else if catFlags.NumberLines {
		result := fmt.Sprintf("%6d\t%s", *lineNumber, line)
		*lineNumber++
		return result
	} else {
		return line
	}
}

// processInput reads from an io.Reader and processes each line
func processInput(reader io.Reader, lineNumber *int) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(formatLine(line, lineNumber))
	}
	return scanner.Err()
}

// processFile opens and processes a single file
func processFile(filename string, lineNumber *int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return processInput(file, lineNumber)
}

var rootCmd = &cobra.Command{
	Use:   "cat-go [flag] [file...]",
	Short: "Concatenate files and print to standard output",
	Long:  `Concatenate and print files to standard output. It can be used to display the content of one or more files. If no file is specified, it reads from standard input.`,
	Run: func(cmd *cobra.Command, args []string) {
		lineNumber := 1
		if len(args) == 0 {
			// Read from standard input
			if err := processInput(os.Stdin, &lineNumber); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			}
			return
		}

		for _, filename := range args {
			if filename == "-" {
				// Read from standard input when "-" is specified
				if err := processInput(os.Stdin, &lineNumber); err != nil {
					fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
				}
				continue
			}

			if err := processFile(filename, &lineNumber); err != nil {
				fmt.Fprintf(os.Stderr, "Error processing file %s: %v\n", filename, err)
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
	rootCmd.Flags().BoolVarP(&catFlags.NumberNonBlank, "number-nonblank", "b", false, "Number non-empty output lines")
	rootCmd.Flags().BoolVarP(&catFlags.NumberLines, "number", "n", false, "Number all output lines")
	rootCmd.Flags().BoolVarP(&catFlags.ShowEnds, "show-ends", "E", false, "Display $ at end of each line")
}
