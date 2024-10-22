package lexer

import (
	"fmt"

	"github.com/jkellogg01/phpls/token"
)

type Lexer struct {
	input   string
	start   int
	current int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.start = l.current
	next := l.advance()
	switch next {
	case '[':
		return l.newToken(token.LSquare)
	case ']':
		return l.newToken(token.RSquare)
	case '(':
		return l.newToken(token.LParen)
	case ')':
		return l.newToken(token.RParen)
	case '{':
		return l.newToken(token.LBrace)
	case '}':
		return l.newToken(token.RBrace)
	case '$':
		return l.newToken(token.Dollar)
	case '\\':
		return l.newToken(token.BSlash)
	case ':':
		return l.newToken(token.Colon)
	case ';':
		return l.newToken(token.Semi)
	case ',':
		return l.newToken(token.Comma)
	default:
		return l.newIllegal("invalid character")
	}
}

func (l *Lexer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: l.input[l.start:l.current],
	}
}

func (l *Lexer) newIllegal(message string) token.Token {
	return token.Token{
		Type:    token.Illegal,
		Literal: fmt.Sprintf("%s at %d", message, l.current),
	}
}

func (l *Lexer) advance() byte {
	if l.current >= len(l.input) {
		return 0
	}
	result := l.input[l.current]
	l.current += 1
	return result
}
