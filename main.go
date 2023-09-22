package main

import (
	"9ccgo/generator"
	"9ccgo/lexer"
	"9ccgo/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough args.")
		return
	}

	tokens := lexer.Tokenize(os.Args[1])
	parser := parser.New(tokens)
	ast := parser.Parse()

	t := generator.Compile(ast)
	fmt.Println(t)
}
