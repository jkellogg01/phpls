package lexer

import (
	"fmt"

	"github.com/jkellogg01/phpls/token"
)

func (l *Lexer) emit(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: l.input[l.start:l.current],
		Row:     l.row,
		Col:     l.col,
	}
}

func (l *Lexer) emitIllegal(message string) token.Token {
	return token.Token{
		Type:    token.Illegal,
		Literal: fmt.Sprintf("%s at %d:%d", message, l.row, l.col),
		Row:     l.row,
		Col:     l.col,
	}
}
