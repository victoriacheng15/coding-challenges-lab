/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// parseURL parses the input URL string and returns a *url.URL
func parseURL(rawURL string) (*url.URL, error) {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}
	return url.Parse(rawURL)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "curl-go [flag][url]",
	Short: "A simple curl-like HTTP client written in Go.",
	Long: `curl-go is a command-line tool for making HTTP requests, similar to curl. 
	It supports GET, POST, PUT, DELETE, custom headers, and data payloads. 
	Use it to interact with web APIs, test endpoints, and inspect responses directly from your terminal.

	Examples:
		curl-go http://example.com
		curl-go -X POST http://example.com -d '{"key":"value"}' -H "Content-Type: application/json"
		curl-go -v http://example.com
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("curl-go: no URL specified!")
			cmd.Println("curl-go: try 'curl-go --help' for more information")
			return
		}
		rawURL := args[0]
		u, err := parseURL(rawURL)
		if err != nil {
			cmd.Println("Invalid URL:", err)
			return
		}
		protocol := u.Scheme
		host := u.Hostname()
		port := u.Port()
		if port == "" && protocol == "http" {
			port = "80"
		}
		path := u.EscapedPath()
		if path == "" {
			path = "/"
		}
		cmd.Println("connecting to", host)
		cmd.Printf("Sending request GET %s HTTP/1.1\n", path)
		cmd.Printf("Host: %s\n", host)
		cmd.Println("Accept: */*")
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.curl-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
