use std::{env, fs};

#[allow(unused_imports)]

fn main() {
    let args: Vec<String> = env::args().collect();

    println!("{:?}", args);
}
