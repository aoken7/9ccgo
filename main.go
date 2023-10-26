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

	fileNmme := os.Args[1]

	source, err := os.ReadFile(fileNmme)
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	tokens := lexer.Tokenize(string(source))
	parser := parser.New(tokens)
	ast := parser.Parse()

	t := generator.Compile(ast)
	fmt.Println(t)
}
