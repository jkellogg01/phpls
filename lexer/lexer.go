package lexer

import (
	"github.com/jkellogg01/phpls/token"
)

type Lexer struct {
	input   string
	start   int
	current int
	row     int
	col     int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:   input,
		start:   0,
		current: 0,
		row:     1,
		col:     0,
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	l.start = l.current
	next := l.advance()
	switch next {
	case 'b', 'B':
		l.advance()
		return l.singleQuotedString()
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
	case '#':
		if l.peek() != '[' {
			return l.singleLineComment()
		}
		l.advance()
		l.newToken(token.PoundLSquare)
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
		switch l.peek() {
		case '/':
			l.advance()
			return l.singleLineComment()
		case '*':
			l.advance()
			return l.delimitedComment()
		case '=':
			l.advance()
			return l.newToken(token.FSlashEq)
		default:
			return l.newToken(token.FSlash)
		}
	case '%':
		return l.check('=', token.PercentEq, token.Percent)
	case '^':
		return l.check('=', token.CaretEq, token.Caret)
	case '?':
		switch l.peek() {
		case '?':
			l.advance()
			return l.check('=', token.TwoQuestionEq, token.TwoQuestion)
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
		switch l.peek() {
		case '=':
			l.advance()
			return l.newToken(token.StarEq)
		case '*':
			l.advance()
			return l.check('=', token.TwoStarEq, token.TwoStar)
		default:
			return l.newToken(token.Star)
		}
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
				// TODO: this should delimit a HereDoc or NowDoc string
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
	case '\'':
		return l.singleQuotedString()
	case 0:
		return l.newToken(token.EOF)
	}
	return l.newIllegal("invalid character")
}
