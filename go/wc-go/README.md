# Coding Challenges - Build Your Own wc Tool

This is a go version of the Unix `wc` command. It allows you to count bytes, lines, words, and characters from a file or input stream. The porject is built with test-driven development in mind.

## What I learned

Through this project, I explored:

- How to design and build a CLI tool using Cobra
- Writing tests first to guide development (TDD)
- Parsing and handling input from both files and streams
- Structuring Go code for maintainability and clarity

## Installation

1. Clone the repository:

```bash
git clone git@github.com:victoriacheng15/coding-challenges-lab.git
```

2. Navigate to the project directory:

```bash
cd go/wc-go
```

3. Download modules

```bash
go mod download

# tidy to ensure go.mod/go.sum are consistent
go mod tidy
```

43. Build

```bash
go build ccwc main.go
```

54. Run the command

```bash
./ccwc test.txt
```

```bash
./ccwc -l test.txt
```

```bash
./ccwc -w test.txt
```

```bash
./ccwc -c test.txt
```


```bash
./ccwc -m test.txt
```

## Links:
- [Coding Challenges Website](https://codingchallenges.fyi/)