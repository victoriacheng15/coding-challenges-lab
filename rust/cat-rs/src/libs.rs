use std::io;

pub fn read_from_stdin() -> Result<String, io::Error> {
    let mut contents = String::new();
    for line in io::stdin().lines() {
        contents.push_str(&line?);
        contents.push_str("\n")
    }

    Ok(contents)
}

pub fn process_input(contents: String, display_line_number: bool, no_number_on_blank: bool) {
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

pub fn display_help() {
    println!("This is a cat command line tool built with Rust!");
    println!();
    println!("cargo run -- [OPTION] [FILE]");
    println!();
    println!("Examples:");
    println!("cargo run -- <file>");
    println!("cargo run -- -n <file>");
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
