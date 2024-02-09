use std::{env, fs, io};

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
        println!("Usage: cargo run -- <file>");
        std::process::exit(1)
    }

    // println!("{:?}", args);

    for filename in args {
        if filename == "-" || filename == "-n" || filename == "-b" {
            let mut contents = String::new();
            for line in io::stdin().lines() {
                contents.push_str(&line?);
                contents.push_str("\n");
            }
            // println!("{filename}");
            process_input(contents, display_line_number, no_number_on_blank)
        } else {
            let contents = fs::read_to_string(filename)?;
            process_input(contents, display_line_number, no_number_on_blank)
        }
    }

    Ok(())
}
