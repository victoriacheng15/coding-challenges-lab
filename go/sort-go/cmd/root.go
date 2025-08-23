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



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sort-go",
	Short: "A brief description of your application",
	Long: `Sort-go is a command-line tool that sorts lines from a specified file using different sorting algorithms, including lexicographical, unique, and random sort. It supports options for removing duplicates, choosing the sorting method, and randomizing output, making it a flexible utility for text processing and experimentation. Designed for extensibility and learning, sort-go helps users understand sorting techniques and command-line application development in Go.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Error: no input file provided\n")
			return
		}

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %q: %v\n", args[0], err)
			return
		}
		defer file.Close()

		// Read lines from file
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading lines: %v\n", err)
			return
		}

		// Sort lines lexicographically
		sort.Strings(lines)

		// Output sorted lines
		for _, line := range lines {
			fmt.Println(line)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sort-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


