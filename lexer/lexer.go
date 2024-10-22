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
	l.skipWhitespace()
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
	case '-':
		switch l.peek() {
		case '>':
			l.advance()
			return l.newToken(token.Arrow)
		case '-':
			l.advance()
			return l.newToken(token.TwoDash)
		case '=':
			l.advance()
			return l.newToken(token.DashEq)
		default:
			return l.newToken(token.Dash)
		}
	case '+':
		switch l.peek() {
		case '+':
			l.advance()
			return l.newToken(token.TwoPlus)
		case '=':
			l.advance()
			return l.newToken(token.PlusEq)
		default:
			return l.newToken(token.Plus)
		}
	case '|':
		switch l.peek() {
		case '|':
			l.advance()
			return l.newToken(token.TwoPipe)
		case '=':
			l.advance()
			return l.newToken(token.PipeEq)
		default:
			return l.newToken(token.Pipe)
		}
	case '&':
		switch l.peek() {
		case '&':
			l.advance()
			return l.newToken(token.TwoAmper)
		case '=':
			l.advance()
			return l.newToken(token.AmperEq)
		default:
			return l.newToken(token.Amper)
		}
	case '/':
		return l.check('=', token.FSlashEq, token.FSlash)
	case '%':
		return l.check('=', token.PercentEq, token.Percent)
	case '^':
		return l.check('=', token.CaretEq, token.Caret)
	case '?':
		switch l.peek() {
		case '?':
			l.advance()
			return l.newToken(token.TwoQuestion)
		case '>':
			l.advance()
			return l.newToken(token.QuestionMore)
		default:
			return l.newToken(token.Question)
		}
	case '=':
		if l.peek() != '=' {
			return l.newToken(token.Eq)
		}
		l.advance()
		return l.check('=', token.ThreeEq, token.TwoEq)
	case '!':
		if l.peek() != '=' {
			return l.newToken(token.Bang)
		}
		l.advance()
		return l.check('=', token.BangTwoEq, token.BangEq)
	case '*':
		if l.peek() == '=' {
			l.advance()
			return l.newToken(token.StarEq)
		} else if l.peek() != '*' {
			return l.newToken(token.Star)
		}
		l.advance()
		return l.check('=', token.TwoStarEq, token.TwoStar)
	case '>':
		if l.peek() == '=' {
			l.advance()
			return l.newToken(token.MoreEq)
		} else if l.peek() != '>' {
			return l.newToken(token.More)
		}
		l.advance()
		return l.check('=', token.TwoMoreEq, token.TwoMore)
	case '.':
		switch l.peek() {
		case '=':
			l.advance()
			return l.newToken(token.DotEq)
		case '.':
			if l.input[l.current+1] != '.' {
				return l.newToken(token.Dot)
			}
			l.advance()
			l.advance()
			return l.newToken(token.ThreeDot)
		default:
			return l.newToken(token.Dot)
		}
	case '<':
		if l.peek() == '=' {
			l.advance()
			return l.check('>', token.LessEqMore, token.LessEq)
		} else if l.peek() == '<' {
			l.advance()
			switch l.peek() {
			case '<':
				l.advance()
				return l.newToken(token.ThreeLess)
			case '=':
				l.advance()
				return l.newToken(token.TwoLessEq)
			default:
				return l.newToken(token.TwoLess)
			}
		} else if l.peek() != '?' {
			return l.newToken(token.Less)
		}
		l.advance()
		next := l.advance()
		if next == '=' {
			return l.newToken(token.EchoOpen)
		}
		// <? and then something other than an equal sign
		for _, c := range "php" {
			if next != byte(c) {
				return l.newIllegal("malformed opening tag")
			}
			next = l.advance()
		}
		return l.newToken(token.Open)
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

func (l *Lexer) skipWhitespace() {
	// TODO: when line numbers are incorporated, this funtion will need to
	// check them - including registering \r\n as only one line break, as if
	// it were a single \r or \n
	for l.peek() == ' ' ||
		l.peek() == '\t' ||
		l.peek() == '\r' ||
		l.peek() == '\n' {
		l.advance()
	}
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
