use std::{env, fs, process};

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

    // println!("{:?}", args);
    Config { options, filename }
}

fn get_field_numbers(options: &[String]) -> Vec<u32> {
    let has_field = options.iter().any(|opt| opt.starts_with("-f"));
    let mut field_number = Vec::new();
    if has_field {
        let field = options
            .iter()
            .find(|opt| opt.starts_with("-f"))
            .unwrap()
            .split("f")
            .nth(1)
            .unwrap();

        for num in field.chars() {
            if num.is_digit(10) {
                field_number.push(num.to_digit(10).unwrap());
            }
        }
    }

    field_number
}

fn get_delimiter(options: &[String]) -> String {
    let _has_delimiter = options.iter().any(|opt| opt.starts_with("-d"));

    //  wait to extract the delimiter for later

    return ",".to_string();
}

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();
    let config = parse_config(&args);

    let field_numbers = get_field_numbers(&config.options);
    let delimiter = get_delimiter(&config.options);

    println!(
        "field_number: {:?}, delimiter: {}",
        field_numbers, delimiter
    );

    let mut all_contents = String::new();
    let contents = fs::read_to_string(&config.filename).expect("Failed to read the file");
    all_contents.push_str(&contents);
    for line in contents.lines() {
        if let Some(column) = line.split_whitespace().nth(1 - 1) {
            println!("{}", column);
        }
    }
}
