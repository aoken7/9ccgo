package lexer

import (
	"9ccgo/token"
	"fmt"
)

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

func (l *Lexer) consumeChar(ch byte) bool {
	if l.peekChar() == ch {
		l.readChar()
		return true
	}
	return false
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.peekChar()) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) readIdentifer() string {
	position := l.position
	if !isDigit(l.ch) && !isLetter(l.ch) {
		panic(fmt.Sprintf("lexer err: %c is not underscore or alphabet.", l.ch))
	}
	for isIdentifer(l.peekChar()) {
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
	case '*':
		tok = newToken(token.ASTERISK, "*")
	case '/':
		tok = newToken(token.SLASH, "/")
	case '<':
		if l.consumeChar('=') {
			tok = newToken(token.LEQ, "<=")
		} else {
			tok = newToken(token.LSS, "<")
		}
	case '>':
		if l.consumeChar('=') {
			tok = newToken(token.GEQ, ">=")
		} else {
			tok = newToken(token.GTR, ">")
		}
	case '=':
		if l.consumeChar('=') {
			tok = newToken(token.EQL, "==")
		} else {
			tok = newToken(token.ASSIGN, "=")
		}
	case '!':
		if l.consumeChar('=') {
			tok = newToken(token.NEQ, "!=")
		}
	case ';':
		tok = newToken(token.SEMICOLON, ";")
	case '(':
		tok = newToken(token.LPAREN, "(")
	case '{':
		tok = newToken(token.LBRACE, "{")
	case ')':
		tok = newToken(token.RPAREN, ")")
	case '}':
		tok = newToken(token.RBRACE, "}")
	default:
		if isDigit(l.ch) {
			num := l.readNumber()
			tok = newToken(token.INT, num)
		} else {
			ident := l.readIdentifer()

			if t, ok := token.Keywords[ident]; ok {
				tok = newToken(t, ident)
			} else {
				tok = newToken(token.IDENT, ident)
			}
		}
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

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isIdentifer(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_'
}
