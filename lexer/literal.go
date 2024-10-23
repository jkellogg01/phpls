package lexer

import (
	"math"

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

func (l *Lexer) hereDocString() token.Token {
	// <<< just ended, at ident start
	identStart := l.current
	var ident string
	for l.current < len(l.input) {
		// NOTE: can't use consumeNewLine here because we need to know the
		// exact index of the first line break character, and consumeNewLine
		// can consume one or two characters
		if l.match('\n') {
			ident = l.input[identStart : l.current-1]
			l.nextRow()
			break
		} else if l.match('\r') {
			ident = l.input[identStart : l.current-1]
			l.match('\n')
			l.nextRow()
			break
		}
		l.advance()
	}
	if l.current >= len(l.input) {
		return l.emitIllegal("unterminated heredoc string")
	}
	if ident == "" {
		return l.emitIllegal("empty heredoc identifier")
	}
	minColumn := math.MaxInt
	for l.current < len(l.input) {
		l.skipWhitespace()
		if l.col < minColumn {
			minColumn = l.col
		}
		c := l.advance()
		if ident[0] != c {
			continue
		}
		i := 1
		for ; i < len(ident); i++ {
			if l.match(ident[i]) {
				continue
			}
			break
		}
		if i < len(ident) {
			break
		}
		if l.col-len(ident) > minColumn {
			return l.emitIllegal("closing heredoc identifier must not be more indented than any part of the heredoc string")
		}
		return l.emit(token.HDString)
	}
	return l.emitIllegal("unterminated heredoc string")
}

func (l *Lexer) nowDocString() token.Token {
	// <<< just ended, at ident start
	l.match('\'')
	identStart := l.current
	var ident string
	for l.current < len(l.input) {
		// NOTE: can't use consumeNewLine here because we need to know the
		// exact index of the first line break character, and consumeNewLine
		// can consume one or two characters
		if l.match('\'') {
			ident = l.input[identStart : l.current-1]
			break
		}
		l.advance()
	}
	if l.current >= len(l.input) {
		return l.emitIllegal("unterminated nowdoc string")
	}
	if ident == "" {
		return l.emitIllegal("empty nowdoc identifier")
	}
	minColumn := math.MaxInt
	for l.current < len(l.input) {
		l.skipWhitespace()
		if l.col < minColumn {
			minColumn = l.col
		}
		c := l.advance()
		if ident[0] != c {
			continue
		}
		i := 1
		for ; i < len(ident); i++ {
			if l.match(ident[i]) {
				continue
			}
			break
		}
		if i < len(ident) {
			break
		}
		if l.col-len(ident) > minColumn {
			return l.emitIllegal("closing nowdoc identifier must not be more indented than any part of the heredoc string")
		}
		return l.emit(token.NDString)
	}
	return l.emitIllegal("unterminated nowdoc string")
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
