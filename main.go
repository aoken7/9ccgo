package main

import (
	"9ccgo/lexer"
	"9ccgo/token"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func compile(s string) string {

	var out bytes.Buffer

	out.WriteString(".intel_syntax noprefix\n")
	out.WriteString(".global main\n")
	out.WriteString("main:\n")

	tokens := lexer.Tokenize(s)

	num, err := strconv.Atoi(tokens.NextToken().Literal)
	if err != nil {
		fmt.Println("Not a number.")
		return ""
	}

	out.WriteString(fmt.Sprintf("\tmov rax, %d\n", num))

	for {
		tok := tokens.NextToken()
		if tok.Type == token.EOF {
			break
		}

		switch tok.Type {
		case token.PLUS:
			operand, err := strconv.Atoi(tokens.NextToken().Literal)
			if err != nil {
				fmt.Println("Not a number.")
				return ""
			}

			out.WriteString(fmt.Sprintf("\tadd rax, %d\n", operand))
		case token.MINUS:
			operand, err := strconv.Atoi(tokens.NextToken().Literal)
			if err != nil {
				fmt.Println("Not a number.")
				return ""
			}

			out.WriteString(fmt.Sprintf("\tsub rax, %d\n", operand))
		}
	}

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
