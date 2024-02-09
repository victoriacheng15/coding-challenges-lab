use std::{env, fs, io};

fn display_help() {
    println!("This is a cat command line tool built with Rust!");
    println!();
    println!("cargo run -- [OPTION] [FILE]");
    println!();
    println!("Examples:");
    println!("cargo run -- <file>");
    println!("head -n3 <file> | cargo run -- -");
    println!("head -n3 <file> | cargo run -- -n");
    println!("sed G <file> | cargo run -- -b");
    println!("sed G <file> | cargo run -- -b | head -n5");
    println!();
    println!("Options:");
    println!();
    println!("-n - display number on all output lines");
    println!();
    println!("-b - display number on nonempty output lines");
    println!();
    println!("--help - display this help");
    println!();
    println!("Thank you for checking this tool out! :D");
    println!();
}

fn read_from_stdin() -> Result<String, io::Error> {
    let mut contents = String::new();
    for line in io::stdin().lines() {
        contents.push_str(&line?);
        contents.push_str("\n")
    }

    Ok(contents)
}

fn process_input(contents: String, display_line_number: bool, no_number_on_blank: bool) {
    let mut line_number = 1;

    for line in contents.lines() {
        let trimmed_line = line.trim();

        if display_line_number {
            println!("{: >6} {}", line_number, trimmed_line);
            line_number += 1
        } else if no_number_on_blank && !trimmed_line.is_empty() {
            println!("{: >6} {}", line_number, trimmed_line);
            line_number += 1
        } else {
            println!("{trimmed_line}")
        }
    }
}

fn main() -> Result<(), io::Error> {
    let args: Vec<String> = env::args().skip(1).collect();
    let display_line_number = args.iter().any(|arg| arg == "-n");
    let no_number_on_blank = args.iter().any(|arg| arg == "-b");

    if args.len() < 1 {
        println!("Usage: cargo run -- [option] <file> <file2>");
        println!("Not sure what to do? type cargo run -- --help");
        std::process::exit(1)
    }

    for filename in args {
        if filename == "-" || filename == "-n" || filename == "-b" {
            let contents = read_from_stdin()?;
            process_input(contents, display_line_number, no_number_on_blank)
        } else if filename == "--help" {
            display_help();
        } else {
            let contents = fs::read_to_string(filename)?;
            process_input(contents, display_line_number, no_number_on_blank)
        }
    }

    Ok(())
}
