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
            "{:?}, [{}:{}] - [{}:{}]",
            self.kind, self.start.0, self.start.1, self.end.0, self.end.1
        )
    }
}

#[derive(Debug)]
pub enum TokenKind {
    Illegal(String),
}
