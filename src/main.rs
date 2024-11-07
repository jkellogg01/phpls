#![allow(dead_code)]

mod lexer;
mod position;
mod token;

use lexer::Lexer;

fn main() {
    println!("Hello, world!");
    let l = Lexer::new("test.txt").unwrap();
    l.for_each(|c| println!("{}", c));
}
