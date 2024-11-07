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

    fn expect(&mut self, expect: u8) -> bool {
        let Some(c) = self.peek() else {
            return false;
        };
        match c == expect {
            true => {
                self.advance();
                true
            }
            false => false,
        }
    }

    fn consume(&mut self, expect: &str, kind: TokenKind) -> Token {
        for x in expect.bytes() {
            if let Some(c) = self.advance() {
                if c == x {
                    continue;
                }
            }
            return self.emit_illegal("unexpected keyword termination");
        }
        self.emit(kind)
    }

    fn emit(&mut self, kind: TokenKind) -> Token {
        let start = self.start.coords();
        let end = self.current.coords();
        self.start = self.current;
        Token::emit(kind, start, end)
    }

    fn emit_illegal(&mut self, message: &str) -> Token {
        self.emit(TokenKind::Illegal(String::from(message)))
    }
}

impl Iterator for Lexer {
    type Item = Token;

    fn next(&mut self) -> Option<Token> {
        match (self.advance(), self.peek()) {
            (Some(b'&'), _) => Some(self.emit(TokenKind::Ampersand)),
            (Some(b'a'), _) => Some(self.consume("bstract", TokenKind::Abstract)),
            (Some(_), _) => Some(self.emit_illegal("unexpected character")),
            (None, _) => None,
        }
    }
}
