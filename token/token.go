package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF   = "EOF"
	IDENT = "IDENT"

	INT = "INT"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	EQL    = "=="
	LSS    = "<"
	GTR    = ">"
	ASSIGN = "="

	NEQ = "!="
	LEQ = "<="
	GEQ = ">="

	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
)
