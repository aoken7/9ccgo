package lexer

import (
	"9ccgo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `1+2;
	12 + 34 - 5;
	1 * 2 - 3;
	(2 + 4) / 3;
	-2;
	1 < 2 == 3 > 1;
	2 <= 3 != 4 >= 5;
	`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "12"},
		{token.PLUS, "+"},
		{token.INT, "34"},
		{token.MINUS, "-"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "1"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
		{token.MINUS, "-"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.SLASH, "/"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.MINUS, "-"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "1"},
		{token.LSS, "<"},
		{token.INT, "2"},
		{token.EQL, "=="},
		{token.INT, "3"},
		{token.GTR, ">"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.LEQ, "<="},
		{token.INT, "3"},
		{token.NEQ, "!="},
		{token.INT, "4"},
		{token.GEQ, ">="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	tokens := Tokenize(input)
	for i, tt := range tests {
		tok := tokens[i]
		if tok.Type != tt.expectType {
			t.Fatalf("got %s, wawnt %s", tok.Type, tt.expectType)
		}
		if tok.Literal != tt.expectLiteral {
			t.Fatalf("got %s, want %s", tok.Literal, tt.expectLiteral)
		}
	}
}

func TestReadNumber(t *testing.T) {
	tests := []string{
		"12345",
		"1",
	}

	for _, tt := range tests {
		tokens := Tokenize(tt)
		tok := tokens[0]
		num := tok.Literal
		if num != tt {
			t.Fatalf("got %s, want %s", num, tt)
		}
	}
}
