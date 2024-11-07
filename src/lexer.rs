use crate::{
    position::Position,
    token::{Token, TokenKind},
};

pub struct Lexer {
    source: Vec<u8>,
    start: Position,
    current: Position,
}

impl Iterator for Lexer {
    type Item = Token;

    fn next(&mut self) -> Option<Token> {
        self.skip_whitespace();
        match (self.advance(), self.peek()) {
            (Some(b'#'), Some(b'[')) => Some(self.consume_attribute()),
            (Some(b'&'), Some(b'&')) => Some(self.consume("&", TokenKind::BooleanAnd)),
            (Some(b'&'), Some(b'=')) => Some(self.emit(TokenKind::AndEqual)),
            (Some(b'&'), _) => Some(self.emit(TokenKind::Ampersand)),
            (Some(b'a'), Some(b'b')) => Some(self.consume("bstract", TokenKind::Abstract)),
            (Some(b'a'), Some(b'r')) => Some(self.consume("rray", TokenKind::Array)),
            (Some(b'a'), Some(b's')) => Some(self.consume("s", TokenKind::As)),
            (Some(b'b'), Some(b'r')) => Some(self.consume("reak", TokenKind::Break)),
            (Some(b'|'), Some(b'|')) => Some(self.consume("|", TokenKind::BooleanOr)),
            (Some(x), _) => {
                if x >= 32 {
                    Some(self.emit_illegal("unexpected character"))
                } else {
                    Some(self.emit(TokenKind::BadCharacter(x)))
                }
            }
            (None, _) => None,
        }
    }
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
            if let Some(c) = self.peek() {
                if c == x {
                    self.advance();
                    continue;
                }
            }
            return self.emit_illegal("unexpected keyword termination");
        }
        self.emit(kind)
    }

    fn consume_attribute(&mut self) -> Token {
        self.expect(b'[');
        let body_start = self.current.index();
        let mut attr_body = String::new();
        while let Some(c) = self.peek() {
            if c != b']' {
                self.advance();
                continue;
            }
            attr_body =
                String::from_utf8_lossy(&self.source[body_start..self.current.index()]).to_string();
            self.expect(b']');
            break;
        }
        self.emit(TokenKind::Attribute(attr_body))
    }

    fn skip_whitespace(&mut self) {
        loop {
            match self.peek() {
                Some(b' ') | Some(b'\t') => {
                    self.advance();
                }
                Some(b'\n') => {
                    self.advance();
                    self.current.row_advance();
                }
                Some(b'\r') => {
                    self.advance();
                    self.expect(b'\n');
                    self.current.row_advance();
                }
                _ => break,
            }
        }
        self.start = self.current;
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
