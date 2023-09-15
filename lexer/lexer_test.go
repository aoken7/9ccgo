package lexer

import "testing"

func TestNextToken(t *testing.T) {
	input := `12+24`

	tests := []string{
		"12",
		"+",
		"24",
		"EOF",
		"EOF",
	}

	lexer := New(input)
	for _, tt := range tests {
		token := lexer.NextToken()
		if token != tt {
			t.Errorf("got %s, want %s", token, tt)
		}
	}
}

func TestReadNumber(t *testing.T) {
	tests := []string{
		"12345",
		"1",
	}

	for _, tt := range tests {
		l := New(tt)
		num := l.readNumber()
		if num != tt {
			t.Fatalf("got %s, want %s", num, tt)
		}
	}
}
