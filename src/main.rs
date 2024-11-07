#![allow(dead_code)]

mod lexer;
mod position;
mod token;

use std::fs;

use lexer::Lexer;

fn main() {
    // TODO: set up both repl and file input
    let Ok(input) = fs::read_to_string("test.txt") else {
        println!("failed to open file!");
        return;
    };
    let l = Lexer::new(input).unwrap();
    l.for_each(|c| println!("{}", c));
}
