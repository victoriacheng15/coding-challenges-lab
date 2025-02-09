import argparse


def validate_field(value):
  try:
    field = int(value)
    if field <= 0:
      raise argparse.ArgumentTypeError(f"Fields are numbered from 1, but got {field}.")
    return field
  except ValueError:
    raise argparse.ArgumentTypeError(f"Invalid field value: {value}. Must be an integer.")

def cut(filename, delimiter, fields):
  try:
    with open(filename, 'r') as f:
      for line in f:
        parts = line.strip().split(delimiter)

        selected_fields = [parts[i-1] for i in fields if i - 1 < len(parts)]

        print(delimiter.join(selected_fields))
  except FileNotFoundError:
    print(f'Error: File "{filename}" not found.')
  except Exception as e:
    print(f'Error: {e}')


def main():
    # create argument parser
    parser = argparse.ArgumentParser(description="A Python implementation of the cut command.")

    # add arguments
    parser.add_argument('filename', type=str, help="The file to process.")
    parser.add_argument('-d', '--delimiter', type=str, default='\t', help="Field delimiter (default is tab).")
    parser.add_argument('-f', '--fields', type=lambda x: [validate_field(f) for f in x.split(',')], required=True, help="Fields to select (e.g., 1,2,3).")

    # parse arguments
    args = parser.parse_args()
    print(args.fields)
    # fields = list(map(int, args.fields.split()))

    # cut(args.filename, args.delimiter, fields)


if __name__ == "__main__":
    main()