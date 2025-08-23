/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"bufio"
	"sort"

	"github.com/spf13/cobra"
)

type Flags struct {
	RemoveDuplicates bool
	SortMethod       string
	RandomSort       bool
}

var sortFlags Flags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sort-go",
	Short: "A brief description of your application",
	Long: `Sort-go is a command-line tool that sorts lines from a specified file using different sorting algorithms, including lexicographical, unique, and random sort. It supports options for removing duplicates, choosing the sorting method, and randomizing output, making it a flexible utility for text processing and experimentation. Designed for extensibility and learning, sort-go helps users understand sorting techniques and command-line application development in Go.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var scanner *bufio.Scanner
		
		if len(args) == 0 {
			// Read from stdin
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			// Read from file
			file, err := os.Open(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %q: %v\n", args[0], err)
				return
			}
			defer file.Close()
			scanner = bufio.NewScanner(file)
		}

		// Read lines from input source
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading lines: %v\n", err)
			return
		}

		if sortFlags.RemoveDuplicates {
			unique := make(map[string]bool)
			var result []string
			for _, line := range lines {
				if !unique[line] {
					unique[line] = true
					result = append(result, line)
				}
			}

			sort.Strings(result)
			lines = result
		} else {
			// Sort lines lexicographically
			sort.Strings(lines)
		}

		// Output sorted lines
		for _, line := range lines {
			fmt.Println(line)
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
	rootCmd.Flags().BoolVarP(&sortFlags.RemoveDuplicates, "remove-duplicates", "u", false, "Remove duplicate lines")
	rootCmd.Flags().StringVarP(&sortFlags.SortMethod, "sort-method", "s", "merge", "Sorting method (merge, quick, binary)")
	rootCmd.Flags().BoolVarP(&sortFlags.RandomSort, "random-sort", "R", false, "Randomize output order")
}


