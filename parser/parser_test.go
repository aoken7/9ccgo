package parser

import (
	"9ccgo/ast"
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
		{"1 * 2", "(1 * 2)"},
		{"1 * 2 / 3", "((1 * 2) / 3)"},
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"(2 + 4) / 3 + (4 - 1) * 2", "(((2 + 4) / 3) + ((4 - 1) * 2))"},
		{"-(1 + 2)", "(-(1 + 2))"},
		{"-3 * 2 + -(1 / 2)", "(((-3) * 2) + (-(1 / 2)))"},
		{"2 < 3 == 1 + 3 > 5", "((2 < 3) == ((1 + 3) > 5))"},
		{"2 <= 3 != 1 + 3 >= 5", "((2 <= 3) != ((1 + 3) >= 5))"},
		{"a = 1", "(a = 1)"},
		{"foo + bar", "(foo + bar)"},
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

func TestExpressionStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1;", "1"},
		{"1 + 2;", "(1 + 2)"},
	}

	for _, tt := range tests {
		tokens := lexer.Tokenize(tt.input)
		p := New(tokens)
		actual := p.Parse()
		compStmt := actual.(*ast.CompoundStatement)
		exprStmt := compStmt.Statements[0].(*ast.ExpressionStatement)
		if exprStmt.String() != tt.expected {
			t.Fatalf("got %s, want %s", exprStmt.String(), tt.expected)
		}
	}
}

func TestCompoundStatement(t *testing.T) {
	input := `
		a = 1;
		b = 2;
		c = 3;
	`

	tokens := lexer.Tokenize(input)
	p := New(tokens)
	actual := p.Parse().String()
	expected := "(a = 1), (b = 2), (c = 3)"
	if actual != expected {
		t.Fatalf("got %s, want %s", actual, expected)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
		return 123;
	`

	tokens := lexer.Tokenize(input)
	p := New(tokens)
	actual := p.Parse().String()
	expected := "123"
	if actual != expected {
		t.Fatalf("got %s, want %s", actual, expected)
	}
}

func TestIfStatement(t *testing.T) {
	input := `
	if (5 == 2 + 3){
		return 5;
	}
	`
	tokens := lexer.Tokenize(input)
	p := New(tokens)
	node := p.Parse()

	cmpStmt, ok := node.(*ast.CompoundStatement)
	if !ok {
		t.Fatalf("node is not ast.Statement. got=%T", node)
	}

	ifStmt, ok := cmpStmt.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("cmpStmt is not *ast.IfStatement. got=%T", ifStmt)
	}

	if ifStmt.Expression.String() != "(5 == (2 + 3))" {
		t.Fatalf("got %s, want %s", ifStmt.Expression.String(), "(5 == (2 + 3))")
	}

	if ifStmt.TrueStatement.String() != "5" {
		t.Fatalf("got %s, want %s", ifStmt.TrueStatement.String(), "5")
	}
}
