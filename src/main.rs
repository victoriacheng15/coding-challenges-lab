use std::env;

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();

    if args.len() < 1 {
        println!("Usage: cargo run -- <file>");
        std::process::exit(1)
    }

    println!("{:?}", args)
}
