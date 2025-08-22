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
		lineNumber := 1
		if len(args) == 0 {
			// Read from standard input
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := scanner.Text()
				originalLine := line
				
				// Add $ at end if showEnds flag is set
				if catFlags.showEnds {
					line += "$"
				}
				
				if catFlags.numberNonBlank {
					if len(originalLine) > 0 { // Check original line for emptiness
						fmt.Printf("%6d\t%s\n", lineNumber, line)
						lineNumber++
					} else {
						fmt.Println(line) // This will just be "$" if showEnds is true
					}
				} else if catFlags.numberLines {
					fmt.Printf("%6d\t%s\n", lineNumber, line)
					lineNumber++
				} else {
					fmt.Println(line)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			}
			return
		}

		for _, filename := range args {
			if filename == "-" {
				// Read from standard input when "-" is specified
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					line := scanner.Text()
					originalLine := line
					
					// Add $ at end if showEnds flag is set
					if catFlags.showEnds {
						line += "$"
					}
					
					if catFlags.numberNonBlank {
						if len(originalLine) > 0 { // Check original line for emptiness
							fmt.Printf("%6d\t%s\n", lineNumber, line)
							lineNumber++
						} else {
							fmt.Println(line) // This will just be "$" if showEnds is true
						}
					} else if catFlags.numberLines {
						fmt.Printf("%6d\t%s\n", lineNumber, line)
						lineNumber++
					} else {
						fmt.Println(line)
					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
				}
				continue
			}

			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
				continue
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				originalLine := line
				
				// Add $ at end if showEnds flag is set
				if catFlags.showEnds {
					line += "$"
				}
				
				if catFlags.numberNonBlank {
					if len(originalLine) > 0 { // Check original line for emptiness
						fmt.Printf("%6d\t%s\n", lineNumber, line)
						lineNumber++
					} else {
						fmt.Println(line) // This will just be "$" if showEnds is true
					}
				} else if catFlags.numberLines {
					fmt.Printf("%6d\t%s\n", lineNumber, line)
					lineNumber++
				} else {
					fmt.Println(line)
				}
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


