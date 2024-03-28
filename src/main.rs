use crate::libs::config::parse_config;
use crate::libs::delimiter::get_delimiter;
use crate::libs::field::get_field_numbers;
use std::{env, fs};

mod libs;

#[allow(unused_imports)]

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();
    let config = parse_config(&args);

    let field_numbers = get_field_numbers(&config.options);
    let delimiter = get_delimiter(&config.options);

    println!(
        "field_number: {:?}, delimiter: {}",
        field_numbers, delimiter
    );
    println!("");

    let contents = fs::read_to_string(&config.filename).expect("Failed to read the file");
    println!("{}", contents);
    println!("");
    for line in contents.lines() {
        if let Some(column) = line.split_whitespace().nth((1 - 1).try_into().unwrap()) {
            println!("{}", column);
        }
    }
}
