use crate::{
    position::Position,
    token::{Token, TokenKind},
};

pub struct Lexer {
    source: Vec<u8>,
    start: Position,
    current: Position,
}

impl Lexer {
    pub fn new(source_string: String) -> std::io::Result<Lexer> {
        return Ok(Lexer {
            source: Vec::from(source_string),
            start: Position::new(),
            current: Position::new(),
        });
    }

    fn advance(&mut self) -> Option<u8> {
        if self.source.len() <= self.current.index() {
            return None;
        }
        let value = self.source[self.current.index()];
        self.current.col_advance();
        return Some(value);
    }

    fn peek(&self) -> Option<u8> {
        if self.source.len() <= self.current.index() {
            return None;
        }
        return Some(self.source[self.current.index()]);
    }

    fn emit(&mut self, kind: TokenKind) -> Token {
        let start = self.start.coords();
        let end = self.current.coords();
        self.start = self.current;
        Token::emit(kind, start, end)
    }
}

impl Iterator for Lexer {
    type Item = Token;

    fn next(&mut self) -> Option<Token> {
        match (self.advance(), self.peek()) {
            (Some(_), _) => {
                Some(self.emit(TokenKind::Illegal(String::from("unexpected character"))))
            }
            (None, _) => None,
        }
    }
}
