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
	a = 2;
	foo + bar;
	return 123;
	if (2 + 3 == 5){ return 10; }
	if (1 > 2) { return 3; } else { return 5;}
	int b;
	int a = 1 + 2;
	int a, b, c = 4;
	int hoge(int a, int b){
		int c = a + b;
		return c;
	}
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
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "foo"},
		{token.PLUS, "+"},
		{token.IDENT, "bar"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.INT, "123"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.EQL, "=="},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.GTR, ">"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.TYPE, "int"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.TYPE, "int"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.TYPE, "int"},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.COMMA, ","},
		{token.IDENT, "c"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.TYPE, "int"},
		{token.IDENT, "hoge"},
		{token.LPAREN, "("},
		{token.TYPE, "int"},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.TYPE, "int"},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.TYPE, "int"},
		{token.IDENT, "c"},
		{token.ASSIGN, "="},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.IDENT, "c"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
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
