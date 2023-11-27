package preprocess

import "testing"

func TestPreprocess(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: `1+2;`, expected: `1+2;`},
		{input: `1+2;//hoge`, expected: `1+2;`},
		{input: `//hoge`, expected: ``},
	}

	for _, tt := range tests {
		actual := Preprocess(tt.input)
		if actual != tt.expected {
			t.Fatalf("got %s, want %s", actual, tt.expected)
		}
	}
}
