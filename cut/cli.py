import sys
import argparse
from cut.core import cut, parse_fields


def main():
    """
    Parses command-line arguments and processes the input (file or stdin).
    """
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
