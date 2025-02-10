# Coding Challenges - Build Your Own cut Tool

This project is a Python implementation of the Unix `cut` command. It allows you to extract specific fields from a file or input stream using a specified delimiter. The project demonstrates the use of `argparse` for handling command-line arguments and `pytest` for writing unit tests.

## What I learned

Through this project, I gained hands-on experience with:

- argparse Module:
  - Learned how to easily set up flags and arguments for a command-line tool.
  - Understood how to handle optional and positional arguments, default values, and help messages.
  - Discovered how argparse simplifies user input validation and error handling.

- Unit Testing with pytest:
  - Wrote unit tests to ensure the core functionality works as expected.
  - Used pytest fixtures like capsys to capture and test output.
  - Learned how to test edge cases and handle exceptions.

- Code Organization:
  - Structured the project into modular components (e.g., core.py for logic, cli.py for the interface).
  - Separated concerns to make the code more maintainable and testable.

- Input/Output Handling:
  - Implemented support for both file input and piped input (stdin).

## Installation

1. Clone the repository:

```bash
git clone https://github.com/victoriacheng15/cc-cut-py.git
```

2. Navigate to the project directory:

```bash
cd cut-command
```

3. Setup a virtual environment (optional but recommended):

```bash
python -m venv env
source env/bin/activate  # On Windows: env\Scripts\activate
```

4. Install dependencies:

```bash
pip install -r requirements.txt
```

## Usage

There are 2 ways to run this:
1. Install and use the `pycut` command:
```bash
pip install -e .
```
2. Run directly without installation:
```bash
python -m cut.cli
```

### Basic Usage
To extract specific fields from a file:
```bash
pycut -d, -f1,2 sample.csv
# or 
python -m cut.cli -d, -f1,2 sample.csv
```

### Piped Input
You can also process input from another command:
```bash
tail -n5 sample.csv | pycut -d, -f1,2
# or
tail -n5 sample.csv | python -m cut.cli -d, -f1,2
```

### Help
To see all available options:
```bash
pycut --help
# or
python -m cut.cli --help
```
 
## Project Structure

```
cut-command/
│
├── cut/                      # Main package
│   ├── __init__.py           # Makes the folder a Python package
│   ├── cli.py                # Command-line interface (argparse logic)
│   ├── core.py               # Core logic (e.g., cut, validate_field, parse_fields)
│
├── tests/                    # Unit tests
│   ├── __init__.py           # Makes tests a package
│   ├── test_core.py          # Tests for core.py
│   ├── test_cli.py           # Tests for cli.py
│
├── requirements.txt          # Dependencies
├── README.md                 # Project documentation
└── setup.py                  # Installation script
```

## Links:
- [Coding Challenges - Build Your Own cut Tool](https://codingchallenges.fyi/challenges/challenge-cut)
- [Coding Challenges Website](https://codingchallenges.fyi/)