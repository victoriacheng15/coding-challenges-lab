import pytest
import argparse
from cut.core import validate_field, parse_fields, cut
from io import StringIO


def test_validate_field():
    """
    Test the validate_field function.
    """
    # Test valid input
    assert validate_field("1") == 1

    # Test invalid input (field <= 0)
    with pytest.raises(argparse.ArgumentTypeError):
        validate_field("0")

    # Test invalid input (non-integer)
    with pytest.raises(argparse.ArgumentTypeError):
        validate_field("abc")


def test_parse_fields():
    """
    Test the parse_fields function.
    """
    # Test valid input (comma-separated)
    assert parse_fields("1,2,3") == [1, 2, 3]

    # Test valid input (space-separated)
    assert parse_fields("1 2 3") == [1, 2, 3]

    # Test invalid input (non-integer)
    with pytest.raises(argparse.ArgumentTypeError):
        parse_fields("1,abc")


def test_cut(capsys):
    """
    Test the cut function.
    """
    # Create a mock input stream
    input_stream = StringIO("a\tb\tc\n1\t2\t3\n")

    # Call the cut function
    cut(input_stream, "\t", [1, 3])

    # Capture the output printed to stdout
    captured = capsys.readouterr()

    # Assert the output matches the expected result
    assert captured.out == "a\tc\n1\t3\n"
