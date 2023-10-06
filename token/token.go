package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF   = "EOF"
	IDENT = "IDENT"

	TYPE = "TYPE"

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
	LBRACE = "{"

	RPAREN = ")"
	RBRACE = "}"

	IF     = "if"
	ELSE   = "else"
	RETURN = "return"
)

var Keywords = map[string]TokenType{
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"int":    TYPE,
}
