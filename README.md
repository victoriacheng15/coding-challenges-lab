# Coding Challenges - Write Your own cat tool

This is a Rust version of the UNIX cat tool. Like the original cat tool, it allows you to:

- Display line number on each line with `-n`
- Display line number on nonempty line with `-b`
- Display the help with `--help`

## Getting Started

1. Installation

```bash
git clone git@github.com:victoriacheng15/cc-cat-rs.git

cd cc-cat-rs
```

2. Have Rust on your computer?

If you have not installed Rust on your computer, please refer to this [Install Rust Guide](https://www.rust-lang.org/tools/install).

3. Once you have Rust installed!

First, run with the cat tool
```bash
cat test.txt
```
Secondly, run with the Rust version of cat tool
```bash
cargo run -- test.txt
```
Lasty, compare the outputs from both in the terminal!

## Usage

|  Flag  | Description                          |
| :----: | :----------------------------------- |
|   -n   | display line number on each line     |
|   -b   | display line number on nonempty line |
| --help | display help                         |

Examples:

```bash
cargo run -- test.txt test2.txt

head -n3 test.txt | cargo run -- -

head -n3 test.txt | cargo run -- -n

sed G test.txt | cargo run -- -b

sed G test.txt | cargo run -- -b | head -n5
```

## What I learned from this?

Initially, I faced a setback where the code couldn't process a single file only (e.g., `cargo run -- test.txt`).

After debugging and investigating, I was able to pinpointed where went wrong. The root cause was the `args` length, I was checking for a minimum of two arguements.

What was the fix? Instead of checking on arguments length where an option may be included, focus on filtering the `args` array to have file name(s). This allows to focus and check on file name(s) solely in the array.

While seeminlky simple, it took time to fix. A good lesson from this: always take breaks! Cannot go wrong with this. When stuck, stepping away clears your head and prevents diving down endless debugging rabbit holes.

## Links

- [Write Your Own cat Tool](https://codingchallenges.fyi/challenges/challenge-cat)
- [Coding Challenges Website](https://codingchallenges.fyi)