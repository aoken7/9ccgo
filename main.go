package main

import (
	"9ccgo/ast"
	"9ccgo/lexer"
	"9ccgo/parser"
	"9ccgo/token"
	"bytes"
	"fmt"
	"os"
)

func calc(node ast.OperatorNode) string {
	var out bytes.Buffer

	out.WriteString("\tpop rdi\n")
	out.WriteString("\tpop rax\n")

	switch node.OperatorType() {
	case token.PLUS:
		out.WriteString("\tadd rax, rdi\n")
	case token.MINUS:
		out.WriteString("\tsub rax, rdi\n")
	case token.ASTERISK:
		out.WriteString("\timul rax, rdi\n")
	case token.SLASH:
		out.WriteString("\tcqo\n")
		out.WriteString("\tidiv, rdi\n")
	}

	out.WriteString("\tpush rax\n")

	return out.String()
}

func gen(node ast.Node) string {
	var out bytes.Buffer

	if n, ok := node.(*ast.PrefixOperatorNode); ok {
		out.WriteString("\tpush 0\n")
		out.WriteString(gen(n.Rhs))
		out.WriteString(calc(n))

		return out.String()
	}

	if n, ok := node.(*ast.IntegerNode); ok {
		return fmt.Sprintf("\tpush %d\n", n.Value)
	}

	n, ok := node.(*ast.InfixOperatorNode)
	if !ok {
		panic("gen error: unexpected node")
	}

	out.WriteString(gen(n.Lhs))
	out.WriteString(gen(n.Rhs))

	out.WriteString(calc(n))

	return out.String()
}

func compile(s string) string {

	var out bytes.Buffer

	out.WriteString(".intel_syntax noprefix\n")
	out.WriteString(".global main\n")
	out.WriteString("main:\n")

	tokens := lexer.Tokenize(s)
	parser := parser.New(tokens)
	ast := parser.Parse()

	out.WriteString(gen(ast))

	out.WriteString("\tpop rax\n")
	out.WriteString("\tret\n")

	return out.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough args.")
		return
	}

	t := compile(os.Args[1])
	fmt.Println(t)
}
