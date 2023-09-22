package parser

import (
	"9ccgo/lexer"
	"testing"
)

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		/* 		{"1 + 2", "(1 + 2)"},
		   		{"1 + 3 - 2", "((1 + 3) - 2)"},
		   		{"1 * 2", "(1 * 2)"},
		   		{"1 * 2 / 3", "((1 * 2) / 3)"},
		   		{"1 + 2 * 3", "(1 + (2 * 3))"},
		   		{"(2 + 4) / 3 + (4 - 1) * 2", "(((2 + 4) / 3) + ((4 - 1) * 2))"},
		   		{"-(1 + 2)", "(-(1 + 2))"},
		   		{"-3 * 2 + -(1 / 2)", "(((-3) * 2) + (-(1 / 2)))"},
		   		{"2 < 3 == 1 + 3 > 5", "((2 < 3) == ((1 + 3) > 5))"},
		   		{"2 <= 3 != 1 + 3 >= 5", "((2 <= 3) != ((1 + 3) >= 5))"}, */
		{"a = 1", "(a = 1)"},
	}

	for _, tt := range tests {
		tokens := lexer.Tokenize(tt.input)
		p := New(tokens)
		actual := p.Parse().String()
		if actual != tt.expected {
			t.Fatalf("got %s, want %s", actual, tt.expected)
		}
	}
}
