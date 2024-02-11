mod libs;
use libs::{display_help, process_input, read_from_stdin};
use std::{env, fs, io, process};

fn main() -> Result<(), io::Error> {
    let args: Vec<String> = env::args().skip(1).collect();
    let display_line_number = args.iter().any(|arg| arg == "-n");
    let no_number_on_blank = args.iter().any(|arg| arg == "-b");
    let help_option = args.iter().any(|arg| arg == "--help");

    if args.len() < 1 {
        println!("Usage: cargo run -- [option] <file> <file2>");
        println!("Not sure what to do? type cargo run -- --help");
        process::exit(1)
    }

    if help_option {
        display_help();
        process::exit(1)
    }

    if args.len() > 1 {
        let filename = args.iter().rev().next().unwrap();
        let contents = fs::read_to_string(filename)?;
        process_input(contents, display_line_number, no_number_on_blank);
    } else {
        let contents = read_from_stdin()?;
        process_input(contents, display_line_number, no_number_on_blank)
    }

    Ok(())
}
