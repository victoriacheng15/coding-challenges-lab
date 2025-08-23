# wc-go

A simple CLI tool written in Go to count lines, words, and bytes in files or standard input, similar to the Unix `wc` command.

## Features
- Count lines, words, and bytes from one or more files, or from standard input
- Display counts for each file and a total when multiple files are provided
- Supports flags to show only lines (`-l`), words (`-w`), or bytes (`-c`)
- Combine flags for custom output

## Installation

Clone the repo and build:

```sh
git clone <repo-url>
cd go/wc-go
go build -o wc-go
```

## Usage

```sh
wc-go [flags] [file...]
```

### Examples

- Count lines, words, and bytes in a file:
  ```sh
  wc-go test.txt
  ```
- Count in multiple files:
  ```sh
  wc-go file1.txt file2.txt
  ```
- Read from standard input:
  ```sh
  echo "hello world" | wc-go
  # or
  echo "hello world" | wc-go -
  ```
- Show only line count:
  ```sh
  wc-go -l test.txt
  ```
- Show only word count:
  ```sh
  wc-go -w test.txt
  ```
- Show only byte count:
  ```sh
  wc-go -c test.txt
  ```
- Combine flags:
  ```sh
  wc-go -lw test.txt
  wc-go -lc test.txt
  wc-go -wc test.txt
  ```

## Links:
- [Coding Challenges Website](https://codingchallenges.fyi/)