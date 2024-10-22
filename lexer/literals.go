package lexer

import (
	"github.com/jkellogg01/phpls/token"
)

func (l *Lexer) singleQuotedString() token.Token {
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
		return l.newIllegal("unterminated string literal")
	}
	// consume closing single quote
	l.advance()
	return l.newToken(token.SQString)
}

func (l *Lexer) singleLineComment() token.Token {
	for l.current < len(l.input) && !l.consumeNewline() {
		if l.peek() != '?' {
			l.advance()
			continue
		}
		if l.current+1 >= len(l.input) ||
			l.input[l.current+1] == '>' {
			break
		}
	}
	return l.newToken(token.Comment)
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
			return l.newToken(token.Comment)
		}
	}
	return l.newIllegal("unterminated delimited comment")
}
