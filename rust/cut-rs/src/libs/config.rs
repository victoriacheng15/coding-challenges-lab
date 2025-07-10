use std::process;

pub struct Config {
    pub options: Vec<String>,
    pub filename: String,
}

pub fn parse_config(args: &[String]) -> Config {
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

    // println!("{:?}", args);
    Config { options, filename }
}
