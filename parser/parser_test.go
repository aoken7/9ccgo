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
		{"1 + 2", "(1 + 2)"},
		{"1 + 3 - 2", "((1 + 3) - 2)"},
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
