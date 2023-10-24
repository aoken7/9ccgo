package parser

import (
	"9ccgo/ast"
	"9ccgo/lexer"
	"9ccgo/types"
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
	}

	for _, tt := range tests {
		tokens := lexer.Tokenize(tt.input)
		p := New(tokens)
		actual := p.compoundStatement(Env{}).String()
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
		compStmt := p.compoundStatement(Env{})
		exprStmt := compStmt.Statements[0].(*ast.ExpressionStatement)
		if exprStmt.String() != tt.expected {
			t.Fatalf("got %s, want %s", exprStmt.String(), tt.expected)
		}
	}
}

func TestCompoundStatement(t *testing.T) {
	input := `
		int a = 1;
		int b = 2;
		int c = 3;
	`

	tokens := lexer.Tokenize(input)
	p := New(tokens)
	actual := p.compoundStatement(Env{env: map[string]int{}}).String()
	expected := "int a = 1; int b = 2; int c = 3;"
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
	actual := p.compoundStatement(Env{}).String()
	expected := "return 123;"
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

	cmpStmt := p.compoundStatement(Env{})

	ifStmt, ok := cmpStmt.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("cmpStmt is not *ast.IfStatement. got=%T", ifStmt)
	}

	if ifStmt.Expression.String() != "(5 == (2 + 3))" {
		t.Fatalf("got %s, want %s", ifStmt.Expression.String(), "(5 == (2 + 3))")
	}

	if ifStmt.TrueStatement.String() != "return 5;" {
		t.Fatalf("got %s, want %s", ifStmt.TrueStatement.String(), "5")
	}
}

func TestIfElseStatement(t *testing.T) {
	input := `
	if (1 > 2){
		return 1;
	} else {
		return 2;
	}
	`
	tokens := lexer.Tokenize(input)
	p := New(tokens)

	cmpStmt := p.compoundStatement(Env{})

	ifStmt, ok := cmpStmt.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("cmpStmt is not *ast.IfStatement. got=%T", ifStmt)
	}

	if ifStmt.Expression.String() != "(1 > 2)" {
		t.Fatalf("got %s, want %s", ifStmt.Expression.String(), "(1 > 2)")
	}

	if ifStmt.TrueStatement.String() != "return 1;" {
		t.Fatalf("got %s, want %s", ifStmt.TrueStatement.String(), "1")
	}

	if ifStmt.FalseStatement.String() != "return 2;" {
		t.Fatalf("got %s, want %s", ifStmt.FalseStatement.String(), "2")
	}
}

func TestDeclaration(t *testing.T) {
	input := `
		int a;
		int b;
		a + b;
		int c;
		1;
	`

	tokens := lexer.Tokenize(input)
	p := New(tokens)
	compStmt := p.compoundStatement(Env{env: map[string]int{}})
	decl := compStmt.Statements[0].(*ast.Declaration)
	if decl.String() != "int a;" {
		t.Fatalf("got %s, want %s", decl.String(), "int a;")
	}
	decl = compStmt.Statements[1].(*ast.Declaration)
	if decl.String() != "int b;" {
		t.Fatalf("got %s, want %s", decl.String(), "int b;")
	}
	exprStmt := compStmt.Statements[2].(*ast.ExpressionStatement)
	if exprStmt.String() != "(a + b)" {
		t.Fatalf("got %s, want %s", exprStmt.String(), "(a + b)")
	}
	decl = compStmt.Statements[3].(*ast.Declaration)
	if decl.String() != "int c;" {
		t.Fatalf("got %s, want %s", decl.String(), "int c;")
	}
}

func TestFunction(t *testing.T) {
	input := `
	int hoge(int a, int b){
		int c = a + b;
		return c;
	}
	`

	tokens := lexer.Tokenize(input)
	p := New(tokens)
	actual := p.Parse()
	fn := actual.(*ast.FunctionNode)
	if fn.Type != types.Int {
		t.Fatalf("declaration-specifier error. got=%s, want=%s", fn.Type, types.Int)
	}
	if fn.Ident.Identifer != "hoge" {
		t.Fatalf("Identifer is missmatch. got=%s, want=%s", fn.Ident.Identifer, "hoge")
	}
	if len(fn.Declarations) != 2 {
		t.Fatalf("Declarations is missmatch. got=%d, want=%d", len(fn.Declarations), 2)
	}
	if len(fn.CmpStmt.Statements) != 2 {
		t.Fatalf("CmpStmt is missmatch. got=%d, want=%d", len(fn.CmpStmt.Statements), 2)
	}
	assign := fn.CmpStmt.Statements[0].(*ast.ExpressionStatement)
	if assign.String() != "int c = (a + b);" {
		t.Fatalf("got %s, want %s", assign.String(), "int c = (a + b);")
	}
	ret := fn.CmpStmt.Statements[1].(*ast.ReturnStatement)
	if ret.String() != "return c;" {
		t.Fatalf("got %s, want %s", ret.String(), "return c;")
	}

}
