package lexer

import (
	"fmt"

	"github.com/jkellogg01/phpls/token"
)

func (l *Lexer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: l.input[l.start:l.current],
		Row:     l.line,
		Col:     l.start,
	}
}

func (l *Lexer) newIllegal(message string) token.Token {
	return token.Token{
		Type:    token.Illegal,
		Literal: fmt.Sprintf("%s at %d", message, l.current),
		Row:     l.line,
		Col:     l.start,
	}
}

func (l *Lexer) skipWhitespace() {
	for (l.peek() == ' ' ||
		l.peek() == '\t' ||
		l.peek() == '\r' ||
		l.peek() == '\n') &&
		l.current < len(l.input) {
		ws := l.advance()
		switch ws {
		case '\r':
			if l.peek() == '\n' {
				l.advance()
			}
			l.line++
		case '\n':
			l.line++
		}
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
	l.line++
	return true
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
	return result
}

func (l *Lexer) peek() byte {
	if l.current >= len(l.input) {
		return 0
	}
	return l.input[l.current]
}
