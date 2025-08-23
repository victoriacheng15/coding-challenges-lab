/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"bufio"
	"sort"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"sort-go/utils"
)

type Flags struct {
	RemoveDuplicates bool
	SortMethod       string
	RandomSort       bool
}

var sortFlags Flags

func removeDuplicates(lines []string) []string {
	unique := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !unique[line] {
			unique[line] = true
			result = append(result, line)
		}
	}
	sort.Strings(result)
	return result
}

func randomize(lines []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	return lines
}

func processContent(args []string) ([]string, error) {
	var scanner *bufio.Scanner
	
	if len(args) == 0 {
		// Read from stdin
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		// Read from file
		file, err := os.Open(args[0]); if err != nil {
			return nil, fmt.Errorf("error opening file %q: %v", args[0], err)
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
		return nil, fmt.Errorf("error reading lines: %v", err)
	}

	return lines, nil
}

var rootCmd = &cobra.Command{
	Use:   "sort-go",
	Short: "A brief description of your application",
	Long: `Sort-go is a command-line tool that sorts lines from a specified file using different sorting algorithms, including lexicographical, unique, and random sort. It supports options for removing duplicates, choosing the sorting method, and randomizing output, making it a flexible utility for text processing and experimentation. Designed for extensibility and learning, sort-go helps users understand sorting techniques and command-line application development in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := processContent(args); if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}

		switch {
			case sortFlags.RemoveDuplicates:
				lines = removeDuplicates(lines)
			case sortFlags.RandomSort:
				lines = randomize(lines)
			case sortFlags.SortMethod == "merge":
				lines = utils.MergeSort(lines)
			case sortFlags.SortMethod == "quick":
				lines = utils.QuickSort(lines)
			case sortFlags.SortMethod == "heap":
				lines = utils.HeapSort(lines)
			default:
				sort.Strings(lines)
		}

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
	rootCmd.Flags().StringVarP(&sortFlags.SortMethod, "sort-method", "s", "merge", "Sorting method (merge, quick, heap)")
	rootCmd.Flags().BoolVarP(&sortFlags.RandomSort, "random-sort", "R", false, "Randomize output order")
}


