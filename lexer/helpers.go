package lexer

import (
	"fmt"

	"github.com/jkellogg01/phpls/token"
)

func (l *Lexer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: l.input[l.start:l.current],
		Row:     l.row,
		Col:     l.col,
	}
}

func (l *Lexer) newIllegal(message string) token.Token {
	return token.Token{
		Type:    token.Illegal,
		Literal: fmt.Sprintf("%s at %d:%d", message, l.row, l.col),
		Row:     l.row,
		Col:     l.col,
	}
}

func (l *Lexer) skipWhitespace() {
	for l.peek() == ' ' ||
		l.peek() == '\t' ||
		l.peek() == '\r' ||
		l.peek() == '\n' {
		if l.current+1 >= len(l.input) {
			break
		}
		if l.consumeNewline() {
			continue
		}
		l.advance()
	}
}

func (l *Lexer) consumeNewline() bool {
	switch l.peek() {
	case '\n':
		l.advance()
	case '\r':
		l.advance()
		if l.peek() == '\n' {
			l.advance()
		}
	default:
		return false
	}
	l.nextRow()
	return true
}

func (l *Lexer) nextRow() {
	l.row += 1
	l.col = 0
}

func (l *Lexer) check(expect byte, present, missing token.TokenType) token.Token {
	if l.peek() != expect {
		return l.newToken(missing)
	}
	l.advance()
	return l.newToken(present)
}

func (l *Lexer) advance() byte {
	if l.current >= len(l.input) {
		return 0
	}
	result := l.input[l.current]
	l.current += 1
	l.col += 1
	return result
}

func (l *Lexer) peek() byte {
	if l.current >= len(l.input) {
		return 0
	}
	return l.input[l.current]
}
