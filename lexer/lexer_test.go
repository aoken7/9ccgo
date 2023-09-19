package lexer

import (
	"9ccgo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `12+24`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.INT, "12"},
		{token.PLUS, "+"},
		{token.INT, "24"},
		{token.EOF, ""},
	}

	tokens := Tokenize(input)
	for _, tt := range tests {
		tok := tokens.NextToken()
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
		tok := tokens.NextToken()
		num := tok.Literal
		if num != tt {
			t.Fatalf("got %s, want %s", num, tt)
		}
	}
}
