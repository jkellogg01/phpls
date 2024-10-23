package lexer

func (l *Lexer) skipWhitespace() {
	for l.peek() == ' ' ||
		l.peek() == '\t' ||
		l.peek() == '\r' ||
		l.peek() == '\n' {
		if l.current >= len(l.input) {
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

func (l *Lexer) match(expect byte) bool {
	if l.peek() != expect {
		return false
	}
	l.advance()
	return true
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

func (l *Lexer) peekNext() byte {
	if l.current+1 >= len(l.input) {
		return 0
	}
	return l.input[l.current+1]
}
