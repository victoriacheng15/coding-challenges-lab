import argparse

def validate_field(value):
    """
    Validates a field value, ensuring it is a positive integer.

    Args:
        value (str): The field value to validate.

    Returns:
        int: The validated field value.

    Raises:
        argparse.ArgumentTypeError: If the field value is not a positive integer.
    """
    try:
        field = int(value)
        if field <= 0:
            raise argparse.ArgumentTypeError(
                f"Fields are numbered from 1, but got {field}."
            )
        return field
    except ValueError:
        raise argparse.ArgumentTypeError(
            f"Invalid field value: {value}. Must be an integer."
        )


def parse_fields(field_input):
    """
    Parses a string of field values, validating each field as a positive integer.

    Args:
        field_input (str): A string of field values, separated by commas or spaces.

    Returns:
        list[int]: A list of validated field values.

    Raises:
        argparse.ArgumentTypeError: If any of the field values are not positive integers.
    """
    normalized_input = field_input.replace(" ", ",")
    fields = normalized_input.split(",")
    return [validate_field(field) for field in fields]


def cut(input_stream, delimiter, fields):
    """
    Cuts the specified fields from each line of the input stream, using the provided delimiter.

    Args:
        input_stream (file-like): The input stream to process.
        delimiter (str): The field delimiter to use.
        fields (list[int]): The field numbers to select (numbered from 1).

    Returns:
        None

    Prints:
        The selected fields for each line, joined by the specified delimiter.
        If a line has fewer fields than requested, only the available fields are printed.
    """
    for line in input_stream:
        parts = line.strip().split(delimiter)
        selected_fields = [parts[i - 1] for i in fields if i - 1 < len(parts)]
        print(delimiter.join(selected_fields))