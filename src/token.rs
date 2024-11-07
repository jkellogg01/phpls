use std::fmt::{Display, Formatter, Result};

pub struct Token {
    kind: TokenKind,
    start: (usize, usize),
    end: (usize, usize),
}

impl Token {
    pub fn emit(kind: TokenKind, start: (usize, usize), end: (usize, usize)) -> Token {
        Token { kind, start, end }
    }
}

impl Display for Token {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(
            f,
            "[{:03}:{:03}]-[{:03}:{:03}] => {:?}",
            self.start.0, self.start.1, self.end.0, self.end.1, self.kind
        )
    }
}

#[derive(Debug)]
pub enum TokenKind {
    Illegal(String),

    Abstract,
    Ampersand,
    AndEqual,
    Array,
    As,
    Attribute(String),
}
