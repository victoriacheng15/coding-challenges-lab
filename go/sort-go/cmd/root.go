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
)

type Flags struct {
	RemoveDuplicates bool
	SortMethod       string
	RandomSort       bool
}

var sortFlags Flags

// Sorting algorithms
func mergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	
	return merge(left, right)
}

func merge(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func quickSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	
	result := make([]string, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []string, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []string, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func heapSort(arr []string) []string {
	result := make([]string, len(arr))
	copy(result, arr)
	n := len(result)
	
	// Build heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}
	
	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}
	
	return result
}

func heapify(arr []string, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

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

		// Remove duplicates if requested
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
		} else if sortFlags.RandomSort {
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(lines), func(i, j int) {
				lines[i], lines[j] = lines[j], lines[i]
			})
		} else if sortFlags.SortMethod == "merge" {
			lines = mergeSort(lines)
		} else if sortFlags.SortMethod == "quick" {
			lines = quickSort(lines)
		} else if sortFlags.SortMethod == "heap" {
			lines = heapSort(lines)
		} else {
			// Default case: lexicographical sort
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
	rootCmd.Flags().StringVarP(&sortFlags.SortMethod, "sort-method", "s", "merge", "Sorting method (merge, quick, heap)")
	rootCmd.Flags().BoolVarP(&sortFlags.RandomSort, "random-sort", "R", false, "Randomize output order")
}


