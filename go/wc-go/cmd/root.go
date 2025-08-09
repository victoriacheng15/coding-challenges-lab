/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Flags struct {
	Lines bool
	Words bool
	Bytes bool
}

var (
	flags Flags
)

var rootCmd = &cobra.Command{
	Use:   "wc-go [flags] [file...]",
	Short: "Count lines, words, and bytes in files or stdin",
	Long:  `wc-go is a simple command-line tool written in Go that counts lines, words, and bytes in input files or standard input. You can pass one or more files as arguments, or pipe input to it. If no flags are provided, wc-go prints all counts.`,
	Run: func(cmd *cobra.Command, args []string) {
		// For now, just print which flags are set and the args
		fmt.Printf("Flags set - Lines: %v, Words: %v, Bytes: %v\n", flags.Lines, flags.Words, flags.Bytes)
		fmt.Printf("Args: %v\n", args)
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
}
