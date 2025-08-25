/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"crypto/tls"
	"io"
	"net"
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
		} else if port == "" && protocol == "https" {
			port = "443"
		}
		path := u.EscapedPath()
		if path == "" {
			path = "/"
		}

		verbose, _ := cmd.Flags().GetBool("verbose")

		// determine host header (include port if nonstandard)
		hostHeader := host
		if (protocol == "http" && port != "80") || (protocol == "https" && port != "443") {
			hostHeader = host + ":" + port
		}

		cmd.Println("connecting to", host)
		if verbose {
			cmd.Printf("> GET %s HTTP/1.1\n", path)
			cmd.Printf("> Host: %s\n", hostHeader)
			cmd.Printf("> Accept: */*\n")
			cmd.Printf("> Connection: close\n")
			cmd.Println(">")
		} else {
			cmd.Printf("Sending request GET %s HTTP/1.1\n", path)
			cmd.Printf("Host: %s\n", hostHeader)
			cmd.Println("Accept: */*")
			cmd.Println("Connection: close")
		}

		// Build the HTTP request string
		request := "GET " + path + " HTTP/1.1\r\n" +
			"Host: " + hostHeader + "\r\n" +
			"Accept: */*\r\n" +
			"Connection: close\r\n" +
			"\r\n"

		// Open TCP connection
		address := net.JoinHostPort(host, port)
		var conn net.Conn
		if protocol == "https" {
			// For HTTPS, use TLS
			tlsConn, err := tls.Dial("tcp", address, nil)
			if err != nil {
				cmd.PrintErrln("Failed to connect:", err)
				return
			}
			conn = tlsConn
		} else {
			tcpConn, err := net.Dial("tcp", address)
			if err != nil {
				cmd.PrintErrln("Failed to connect:", err)
				return
			}
			conn = tcpConn
		}
		defer conn.Close()

		// Send request
		_, err = conn.Write([]byte(request))
		if err != nil {
			cmd.PrintErrln("Failed to send request:", err)
			return
		}

		// Read response: parse headers then stream body
		reader := bufio.NewReader(conn)

		// Status line
		statusLine, err := reader.ReadString('\n')
		if err != nil {
			cmd.PrintErrln("Failed to read response:", err)
			return
		}
		if verbose {
			cmd.Printf("< %s", statusLine)
		} else {
			os.Stdout.WriteString(statusLine)
		}

		// Headers
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				cmd.PrintErrln("Failed to read headers:", err)
				return
			}
			if line == "\r\n" || line == "\n" {
				// end of headers
				if verbose {
					cmd.Println("<")
				} else {
					os.Stdout.WriteString("\n")
				}
				break
			}
			if verbose {
				// print header lines with '< ' prefix
				// Use Printf to avoid double newlines
				cmd.Printf("< %s", line)
			} else {
				os.Stdout.WriteString(line)
			}
		}

		// Body: stream remaining data
		_, err = io.Copy(os.Stdout, reader)
		if err != nil && err != io.EOF {
			cmd.PrintErrln("Error reading body:", err)
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
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
}
