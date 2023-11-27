package main

import (
	"9ccgo/generator"
	"9ccgo/lexer"
	"9ccgo/parser"
	"9ccgo/preprocess"
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

	input := preprocess.Preprocess(string(source))
	tokens := lexer.Tokenize(input)
	parser := parser.New(tokens)
	ast := parser.Parse()

	t := generator.Compile(ast)
	fmt.Println(t)
}
