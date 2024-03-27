use std::{env, fs, io, process};

#[allow(unused_imports)]

struct Config {
    options: Vec<String>,
    filename: String,
}

fn parse_config(args: &[String]) -> Config {
    let options: Vec<String> = args
        .iter()
        .filter(|arg| arg.contains("-"))
        .cloned()
        .collect();
    let filename: String = args
        .iter()
        .filter(|arg| !arg.contains("-"))
        .cloned()
        .collect();

    if args.is_empty() {
        println!("Usage: cargo run -- [options] <file>");
        // println!("Not sure what to do? type cargo run -- --help");
        process::exit(1);
    }

    println!("{:?}", args);
    Config { options, filename }
}

fn main() -> Result<(), io::Error> {
    let args: Vec<String> = env::args().skip(1).collect();
    let config = parse_config(&args);

    let mut all_contents = String::new();

    let contents = fs::read_to_string(&config.filename)?;
    all_contents.push_str(&contents);

    println!("{}", contents);
    println!("{:?}", config.options);

    Ok(())
}
