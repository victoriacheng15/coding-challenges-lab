# sort-go

A flexible CLI tool written in Go to sort lines from files or standard input, similar to the Unix `sort` command, with support for multiple sorting algorithms and options.

## Features
- Sort lines from one or more files, or from standard input
- Remove duplicate lines (`-u`)
- Choose sorting algorithm: merge, quick, or heap (`-s`)
- Randomize output order (`-R`)
- Combine flags for custom output

## Installation

Clone the repo and build:

```sh
git clone git@github.com:victoriacheng15/coding-challenges-lab.git
cd go/sort-go
go build -o sort-go
```

## Usage

```sh
sort-go [flags] [file...]
```

### Examples

- Sort contents of a file lexicographically:
  ```sh
  sort-go test.txt
  ```
- Sort contents of multiple files:
  ```sh
  sort-go file1.txt file2.txt
  ```
- Read from standard input:
  ```sh
  echo -e "c\nb\na" | sort-go
  ```
- Remove duplicate lines:
  ```sh
  sort-go -u test.txt
  ```
- Use quick sort algorithm:
  ```sh
  sort-go -s quick test.txt
  ```
- Use heap sort algorithm:
  ```sh
  sort-go -s heap test.txt
  ```
- Randomize output order:
  ```sh
  sort-go -R test.txt
  ```
- Combine flags:
  ```sh
  sort-go -u -s quick test.txt
  sort-go -R -s heap test.txt
  ```
