package lexer

import "9ccgo/token"

type Lexer struct {
	input        string
	position     int //今見てるトークンの位置
	readPosition int
	ch           byte
}

func newLexer(input string) *Lexer {
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

func (l *Lexer) eatSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
}

func newToken(t token.TokenType, s string) token.Token {
	return token.Token{Type: t, Literal: s}
}

func (l *Lexer) nextToken() token.Token {

	l.eatSpace()

	if l.readPosition > len(l.input) {
		return newToken(token.EOF, "")
	}

	var tok token.Token

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, "+")
	case '-':
		tok = newToken(token.MINUS, "-")
	case ';':
		tok = newToken(token.SEMICOLON, ";")
	default:
		num := l.readNumber()
		tok = newToken(token.INT, num)
	}

	l.readChar()
	return tok
}

func Tokenize(input string) []token.Token {
	tokenes := make([]token.Token, 0)

	lexer := newLexer(input)
	for {
		tok := lexer.nextToken()
		tokenes = append(tokenes, tok)
		if tok.Type == token.EOF {
			break
		}
	}

	return tokenes
}
