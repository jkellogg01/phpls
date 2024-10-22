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
