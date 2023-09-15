package lexer

type Lexer struct {
	input        string
	position     int //今見てるトークンの位置
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for l.isDigit(l.peekChar()) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) NextToken() string {
	if l.readPosition > len(l.input) {
		return "EOF"
	}

	tok := ""

	switch l.ch {
	case '+':
		tok = "+"
	case '-':
		tok = "-"
	default:
		tok = l.readNumber()
	}

	l.readChar()
	return tok
}
