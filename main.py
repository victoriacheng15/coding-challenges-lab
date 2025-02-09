import sys
import argparse


def validate_field(value):
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
    normalized_input = field_input.replace(" ", ",")
    fields = normalized_input.split(",")
    return [validate_field(field) for field in fields]


def cut(input_stream, delimiter, fields):
    for line in input_stream:
        parts = line.strip().split(delimiter)
        selected_fields = [parts[i - 1] for i in fields if i - 1 < len(parts)]
        print(delimiter.join(selected_fields))


def main():
    # create argument parser
    parser = argparse.ArgumentParser(
        description="A Python implementation of the cut command."
    )

    # add arguments
    parser.add_argument("filename", type=str, nargs="?", help="The file to process.")
    parser.add_argument(
        "-d",
        "--delimiter",
        type=str,
        default="\t",
        help="Field delimiter (default is tab).",
    )
    parser.add_argument(
        "-f",
        "--fields",
        type=parse_fields,
        required=True,
        help="Fields to select (e.g., 1,2,3).",
    )

    # parse arguments
    args = parser.parse_args()

    if args.filename:
        with open(args.filename, "r") as file:
            cut(file, args.delimiter, args.fields)
    else:
        cut(sys.stdin, args.delimiter, args.fields)


if __name__ == "__main__":
    main()
