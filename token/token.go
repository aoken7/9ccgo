package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = "EOF"

	INT = "INT"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	EQL = "=="
	LSS = "<"
	GTR = ">"

	NEQ = "!="
	LEQ = "<="
	GEQ = ">="

	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
)
