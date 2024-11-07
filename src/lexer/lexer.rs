use crate::lexer::token::Token;
use std::fs;

pub struct Lexer {
    source: Vec<u8>,
    start: usize,
    current: usize,
}

impl Lexer {
    pub fn new(source_path: &str) -> std::io::Result<Lexer> {
        let source_bytes = fs::read(source_path)?;
        return Ok(Lexer {
            source: source_bytes,
            start: 0,
            current: 0,
        });
    }

    fn advance(&mut self) -> u8 {
        self.start = self.current;
        if self.source.len() <= self.current {
            return b'\0';
        }
        self.current += 1;
        return self.source[self.start];
    }
}

impl Iterator for Lexer {
    type Item = Token;

    fn next(&mut self) -> Option<Token> {
        None
    }
}
