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
        match (self.advance(), self.peek(), self.peek_next()) {
            (b'#', b'[', _) => Some(self.consume_attribute()),
            (b'#', _, _) => Some(self.consume_comment_line()),
            (b'%', b'>', _) | (b'?', b'>', _) => Some(self.consume(">", TokenKind::CloseTag)),
            (b'&', b'&', _) => Some(self.consume("&", TokenKind::BooleanAnd)),
            (b'&', b'=', _) => Some(self.emit(TokenKind::AndEqual)),
            (b'&', _, _) => Some(self.emit(TokenKind::Ampersand)),
            (b'.', b'=', _) => Some(self.consume("=", TokenKind::ConcatEqual)),
            (b'/', b'*', _) => Some(self.consume_comment_delim()),
            (b'/', b'/', _) => Some(self.consume_comment_line()),
            (b'?', b'?', b'=') => Some(self.consume("?=", TokenKind::CoalesceEqual)),
            (b'?', b'?', _) => Some(self.consume("?", TokenKind::Coalesce)),
            (b'_', b'_', b'C') => Some(self.consume("_CLASS__", TokenKind::ClassConst)),
            (b'a', b'b', b's') => Some(self.consume("bstract", TokenKind::Abstract)),
            (b'a', b'r', b'r') => Some(self.consume("rray", TokenKind::Array)),
            (b'a', b's', _) => Some(self.consume("s", TokenKind::As)),
            (b'b', b'r', b'e') => Some(self.consume("reak", TokenKind::Break)),
            (b'c', b'a', b'l') => Some(self.consume("allable", TokenKind::Callable)),
            (b'c', b'a', b's') => Some(self.consume("ase", TokenKind::Case)),
            (b'c', b'a', b't') => Some(self.consume("atch", TokenKind::Catch)),
            (b'c', b'l', b'a') => Some(self.consume("lass", TokenKind::Class)),
            (b'c', b'l', b'o') => Some(self.consume("lone", TokenKind::Clone)),
            (b'c', b'o', b'n') => match self.peek_at(2) {
                b's' => Some(self.consume("onst", TokenKind::Const)),
                b't' => Some(self.consume("ontinue", TokenKind::Continue)),
                _ => Some(self.emit_illegal("unexpected character")),
            },
            (b'|', b'|', _) => Some(self.consume("|", TokenKind::BooleanOr)),
            (b'\0', _, _) => None,
            (x, _, _) => {
                if x >= 32 {
                    Some(self.emit_illegal("unexpected character"))
                } else {
                    Some(self.emit(TokenKind::BadCharacter(x)))
                }
            }
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

    fn advance(&mut self) -> u8 {
        let at = self.current.index();
        if at > self.source.len() {
            return b'\0';
        } else {
            self.current.col_advance();
            self.source[at]
        }
    }

    fn peek(&self) -> u8 {
        self.peek_at(0)
    }

    fn peek_next(&self) -> u8 {
        self.peek_at(1)
    }

    fn peek_at(&self, offset: usize) -> u8 {
        let at = self.current.index() + offset;
        if at >= self.source.len() {
            b'\0'
        } else {
            self.source[at]
        }
    }

    fn expect(&mut self, expect: u8) -> bool {
        match self.peek() == expect {
            true => {
                self.advance();
                true
            }
            false => false,
        }
    }

    fn consume(&mut self, expect: &str, kind: TokenKind) -> Token {
        for x in expect.bytes() {
            if self.peek() == x {
                self.advance();
                continue;
            }
            return self.emit_illegal("unexpected keyword termination");
        }
        self.emit(kind)
    }

    fn consume_attribute(&mut self) -> Token {
        self.expect(b'[');
        let body_start = self.current.index();
        loop {
            let c = self.peek();
            if c == b'\0' {
                return self.emit_illegal("unterminated attribute");
            } else if c != b']' {
                self.advance();
                continue;
            }
            let attr_body =
                String::from_utf8_lossy(&self.source[body_start..self.current.index()]).to_string();
            self.expect(b']');
            return self.emit(TokenKind::Attribute(attr_body));
        }
    }

    fn consume_comment_line(&mut self) -> Token {
        loop {
            match self.advance() {
                b'\0' => break,
                b'\n' => {
                    self.current.row_advance();
                    break;
                }
                b'\r' => {
                    self.expect(b'\n');
                    self.current.row_advance();
                    break;
                }
                _ => continue,
            };
        }
        self.emit(TokenKind::Comment)
    }

    fn consume_comment_delim(&mut self) -> Token {
        loop {
            let c = self.advance();
            if c == b'\0' {
                return self.emit_illegal("unterminated comment");
            } else if c != b'*' {
                continue;
            }
            if self.expect(b'/') {
                return self.emit(TokenKind::Comment);
            }
        }
    }

    fn skip_whitespace(&mut self) {
        loop {
            match self.peek() {
                b' ' | b'\t' => {
                    self.advance();
                }
                b'\n' => {
                    self.advance();
                    self.current.row_advance();
                }
                b'\r' => {
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
