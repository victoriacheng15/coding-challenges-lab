# cat-go

A simple CLI tool written in Go to concatenate files and print to standard output, similar to the Unix `cat` command.

## Features
- Concatenate and display contents of one or more files
- Read from standard input if no files are specified or if `-` is used as a filename
- Number all lines (`-n`)
- Number non-blank lines only (`-b`)
- Show `$` at the end of each line (`-E`)
- Combine flags for custom output

## Installation

Clone the repo and build:

```sh
git clone <repo-url>
cd go/cat-go
go build -o cat-go
```

## Usage

```sh
cat-go [flags] [file...]
```

### Examples

- Print contents of a file:
  ```sh
  cat-go test.txt
  ```
- Print contents of multiple files:
  ```sh
  cat-go file1.txt file2.txt
  ```
- Read from standard input:
  ```sh
  echo "hello" | cat-go
  # or 
  echo "hello" | cat-go -
  ```
- Number all lines:
  ```sh
  cat-go -n test.txt
  ```
- Number non-blank lines:
  ```sh
  cat-go -b test.txt
  ```
- Show ends:
  ```sh
  cat-go -E test.txt
  ```
- Combine flags:
  ```sh
  cat-go -nE test.txt
  cat-go -bE test.txt
  ```