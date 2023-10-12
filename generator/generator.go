package generator

import (
	"9ccgo/ast"
	"9ccgo/token"
	"bytes"
	"fmt"
	"strconv"
)

var jumpLabel = -1

func getJumpLabel() string {
	jumpLabel++
	return ".Lend" + strconv.Itoa(jumpLabel)
}

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

func genLval(node ast.IdentiferNode) string {
	var out bytes.Buffer

	out.WriteString("\tmov rax, rbp\n")
	out.WriteString(fmt.Sprintf("\tsub rax, %d\n", node.Offset))
	out.WriteString("\tpush rax\n")

	return out.String()
}

func gen(node ast.Node) string {
	var out bytes.Buffer

	switch n := node.(type) {
	case *ast.CompoundStatement:
		for _, stmt := range n.Statements {
			out.WriteString(gen(stmt))
			out.WriteString("\tpop rax\n")
		}
		return out.String()

	case *ast.ExpressionStatement:
		return gen(n.Expression)

	case *ast.ReturnStatement:
		out.WriteString(gen(n.Expression))
		out.WriteString("\tpop rax\n")
		out.WriteString("\tmov rsp, rbp\n")
		out.WriteString("\tpop rbp\n")
		out.WriteString("\tret\n")
		return out.String()

	case *ast.IfStatement:
		out.WriteString(gen(n.Expression))
		out.WriteString("\tpop rax\n")
		out.WriteString("\tcmp rax, 0\n")

		label := getJumpLabel()

		out.WriteString("\tje " + label + "\n")
		out.WriteString(gen(n.TrueStatement))
		out.WriteString(label + ":\n")

		if n.FalseStatement != nil {
			out.WriteString(gen(n.FalseStatement))
		}

		return out.String()

	case *ast.PrefixOperatorNode:
		out.WriteString("\tpush 0\n")
		out.WriteString(gen(n.Rhs))
		out.WriteString(calc(n))
		return out.String()

	case *ast.IntegerNode:
		return fmt.Sprintf("\tpush %d\n", n.Value)

	case *ast.IdentiferNode:
		out.WriteString(genLval(*n))
		out.WriteString("\tpop rax\n")
		out.WriteString("\tmov rax, [rax]\n")
		out.WriteString("\tpush rax\n")
		return out.String()

	case *ast.Declaration:
		out.WriteString(genLval(n.Ident))
		if n.Right != nil {
			out.WriteString(gen(n.Right))
			out.WriteString("\tpop rdi\n")
			out.WriteString("\tpop rax\n")
			out.WriteString("\tmov [rax], rdi\n")
			out.WriteString("\tpush rdi\n")
		}
		out.WriteString("\tsub rsp, 8\n")
		return out.String()
	}

	n, ok := node.(*ast.InfixOperatorNode)
	if !ok {
		panic(fmt.Sprintf("gen error: node is not *ast.InfixOperatorNode. got=%T.", node))
	}

	out.WriteString(gen(n.Lhs))
	out.WriteString(gen(n.Rhs))

	out.WriteString(calc(n))

	return out.String()
}

func Compile(node ast.Node) string {

	var out bytes.Buffer

	out.WriteString(".intel_syntax noprefix\n")
	out.WriteString(".global main\n")
	out.WriteString("main:\n")

	out.WriteString("\tpush rbp\n")
	out.WriteString("\tmov rbp, rsp\n")

	out.WriteString(gen(node))

	out.WriteString("\tmov rsp, rbp\n")
	out.WriteString("\tpop rbp\n")
	out.WriteString("\tret\n")

	return out.String()
}
