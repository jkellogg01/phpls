package lexer

import (
	"github.com/jkellogg01/phpls/token"
)

func (l *Lexer) singleQuoteString() token.Token {
	// at this point, we have _just_ consumed the opening single quote
	for l.peek() != '\'' && l.current < len(l.input) {
		if l.consumeNewline() {
			continue
		}
		if l.peek() == '\\' {
			l.advance()
			if l.consumeNewline() {
				continue
			}
		}
		l.advance()
	}
	if l.current >= len(l.input) {
		return l.emitIllegal("unterminated string literal")
	}
	// consume closing single quote
	l.advance()
	return l.emit(token.SQString)
}

func (l *Lexer) doubleQuoteString() token.Token {
	for l.peek() != '"' && l.current < len(l.input) {
		if l.consumeNewline() {
			continue
		}
		if l.peek() == '\\' {
			l.advance()
			if l.consumeNewline() {
				continue
			}
		}
		l.advance()
	}
	if l.current >= len(l.input) {
		return l.emitIllegal("unterminated string literal")
	}
	// consume closing double quote
	l.advance()
	return l.emit(token.DQString)
}

func (l *Lexer) singleLineComment() token.Token {
	for l.current < len(l.input) && !l.consumeNewline() {
		if l.peek() != '?' {
			l.advance()
			continue
		}
		if l.peekNext() == '>' {
			break
		}
	}
	return l.emit(token.Comment)
}

func (l *Lexer) delimitedComment() token.Token {
	for l.current < len(l.input) {
		c := l.advance()
		if c != '*' {
			continue
		}
		if l.current >= len(l.input) {
			break
		}
		c = l.advance()
		if c == '/' {
			return l.emit(token.Comment)
		}
	}
	return l.emitIllegal("unterminated delimited comment")
}
