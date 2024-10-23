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
		if l.match('\'') {
			return l.singleQuoteString()
		} else if l.match('"') {
			return l.doubleQuoteString()
		}
	case '[':
		return l.emit(token.LSquare)
	case ']':
		return l.emit(token.RSquare)
	case '(':
		return l.emit(token.LParen)
	case ')':
		return l.emit(token.RParen)
	case '{':
		return l.emit(token.LBrace)
	case '}':
		return l.emit(token.RBrace)
	case '#':
		if l.match('[') {
			return l.emit(token.PoundLSquare)
		}
		return l.singleLineComment()
	case '$':
		return l.emit(token.Dollar)
	case '\\':
		return l.emit(token.BSlash)
	case ':':
		return l.emit(token.Colon)
	case ';':
		return l.emit(token.Semi)
	case ',':
		return l.emit(token.Comma)
	case '-':
		if l.match('>') {
			return l.emit(token.Arrow)
		} else if l.match('-') {
			return l.emit(token.TwoDash)
		} else if l.match('=') {
			return l.emit(token.DashEq)
		}
		return l.emit(token.Dash)
	case '+':
		if l.match('+') {
			return l.emit(token.TwoPlus)
		} else if l.match('=') {
			return l.emit(token.PlusEq)
		}
		return l.emit(token.Plus)
	case '|':
		if l.match('|') {
			return l.emit(token.TwoPipe)
		} else if l.match('=') {
			return l.emit(token.PipeEq)
		}
		return l.emit(token.Pipe)
	case '&':
		if l.match('&') {
			return l.emit(token.TwoAmper)
		} else if l.match('=') {
			return l.emit(token.AmperEq)
		}
		return l.emit(token.Amper)
	case '/':
		if l.match('/') {
			return l.singleLineComment()
		} else if l.match('*') {
			return l.delimitedComment()
		} else if l.match('=') {
			return l.emit(token.FSlashEq)
		}
		return l.emit(token.FSlash)
	case '%':
		if l.match('=') {
			return l.emit(token.PercentEq)
		}
		return l.emit(token.Percent)
	case '^':
		if l.match('=') {
			return l.emit(token.CaretEq)
		}
		return l.emit(token.Caret)
	case '?':
		if l.match('?') {
			if l.match('=') {
				return l.emit(token.TwoQuestionEq)
			}
			return l.emit(token.TwoQuestion)
		} else if l.match('>') {
			return l.emit(token.QuestionMore)
		}
		return l.emit(token.Question)
	case '=':
		if !l.match('=') {
			return l.emit(token.Eq)
		}
		if l.match('=') {
			return l.emit(token.ThreeEq)
		}
		return l.emit(token.TwoEq)
	case '!':
		if !l.match('=') {
			return l.emit(token.Bang)
		}
		if l.match('=') {
			return l.emit(token.BangTwoEq)
		}
		return l.emit(token.BangEq)
	case '*':
		if l.match('=') {
			return l.emit(token.StarEq)
		} else if l.match('*') {
			if l.match('=') {
				return l.emit(token.TwoStarEq)
			}
			return l.emit(token.TwoStar)
		}
		return l.emit(token.Star)
	case '>':
		if l.match('=') {
			return l.emit(token.MoreEq)
		} else if !l.match('>') {
			return l.emit(token.More)
		}
		if l.match('=') {
			return l.emit(token.TwoMoreEq)
		}
		return l.emit(token.TwoMore)
	case '.':
		if l.match('=') {
			return l.emit(token.DotEq)
		} else if l.peek() != '.' {
			return l.emit(token.Dot)
		}
		if l.peekNext() != '.' {
			return l.emit(token.Dot)
		}
		l.advance()
		l.advance()
		return l.emit(token.ThreeDot)
	case '<':
		if l.match('=') {
			if l.match('>') {
				return l.emit(token.LessEqMore)
			}
			return l.emit(token.LessEq)
		} else if l.match('<') {
			if l.match('<') {
				if l.match('\'') {
					return l.nowDocString()
				}
				return l.hereDocString()
			} else if l.match('=') {
				return l.emit(token.TwoLessEq)
			}
			return l.emit(token.TwoLess)
		} else if !l.match('?') {
			return l.emit(token.Less)
		}
		if l.match('=') {
			return l.emit(token.EchoOpen)
		}
		// <? and then something other than an equal sign
		for _, c := range "php" {
			if l.match(byte(c)) {
				continue
			}
			return l.emitIllegal("malformed opening tag")
		}
		return l.emit(token.Open)
	case '\'':
		return l.singleQuoteString()
	case '"':
		return l.doubleQuoteString()
	case 0:
		return l.emit(token.EOF)
	}
	return l.emitIllegal("invalid character")
}
