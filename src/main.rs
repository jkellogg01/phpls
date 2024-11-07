mod lexer;

use lexer::{Lexer, Token};

fn main() {
    println!("Hello, world!");
    let l = Lexer::new("test.txt").unwrap();
    l.for_each(|c| println!("{}", c));

    let t = Token::emit();
    println!("{}", t);
}
