# Coding Challenges - Write Your own cat tool

This is a Rust version of the UNIX cat tool. Like the original cat tool, it allows you to:

- Display line number on each line with `-n`
- Display line number on nonempty line with `-b`
- Display the help with `--help`

## Getting Started

1. Installation

```bash
git clone git@github.com:victoriacheng15/coding-challenges-lab.git

cd rust/cat-rs
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

## What did I learn from this?

Initially, I encountered a setback where the code wouldn't process a single file only (e.g., `cargo run -- test.txt`).

After debugging with logging, I zeroed in on the culprit: the `args` length. I had been checking for a minimum of two arguments. This was the reason why the program wouldn’t process a single file.

What was the fix? To effectively read contents from files, we need the files themselves. Thereforce, I ensured the `args` array exclusively contains file paths. This way, I can directly determine the number of files provided. If files exist, the program retrieves their contents; otherwise, it reads from standard input.

Key Takeaways:
- Take breaks to refresh your perspective. Stepping away can help you see things differently and identify solutions.
- Utilize a process of elimination to isolate the root cause. Systematically rule out possibilities until you pinpoint the issue.


## Links

- [Write Your Own cat Tool](https://codingchallenges.fyi/challenges/challenge-cat)
- [Coding Challenges Website](https://codingchallenges.fyi)
