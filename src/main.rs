use std::{env, fs, io};

fn main() -> Result<(), io::Error> {
    let args: Vec<String> = env::args().skip(1).collect();

    if args.len() < 1 {
        println!("Usage: cargo run -- <file>");
        std::process::exit(1)
    }

    println!("{:?}", args);

    for filename in args {
        if filename == "" {
            println!("no action yet")
        } else {
            let contents = fs::read_to_string(filename)?;
            println!("{contents}")
        }
    }

    Ok(())
}
