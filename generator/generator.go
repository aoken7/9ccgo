package generator

import (
	"9ccgo/ast"
	"9ccgo/token"
	"bytes"
	"fmt"
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
	case token.EQL:
		out.WriteString("\tcmp rax, rdi\n")
		out.WriteString("\tsete al\n")
		out.WriteString("\tmovzb rax, al\n")
	case token.NEQ:
		out.WriteString("\tcmp rax, rdi\n")
		out.WriteString("\tsetne al\n")
		out.WriteString("\tmovzb rax, al\n")
	case token.LSS:
		out.WriteString("\tcmp rax, rdi\n")
		out.WriteString("\tsetl al\n")
		out.WriteString("\tmovzb rax, al\n")
	case token.LEQ:
		out.WriteString("\tcmp rax, rdi\n")
		out.WriteString("\tsetle al\n")
		out.WriteString("\tmovzb rax, al\n")
	case token.GTR:
		out.WriteString("\tcmp rdi, rax\n")
		out.WriteString("\tsetl al\n")
		out.WriteString("\tmovzb rax, al\n")
	case token.GEQ:
		out.WriteString("\tcmp rdi, rax\n")
		out.WriteString("\tsetle al\n")
		out.WriteString("\tmovzb rax, al\n")
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

func Compile(ast ast.Node) string {

	var out bytes.Buffer

	out.WriteString(".intel_syntax noprefix\n")
	out.WriteString(".global main\n")
	out.WriteString("main:\n")

	out.WriteString(gen(ast))

	out.WriteString("\tpop rax\n")
	out.WriteString("\tret\n")

	return out.String()
}
