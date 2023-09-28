package main

import (
	"9ccgo/generator"
	"9ccgo/lexer"
	"9ccgo/parser"
	"syscall/js"
)

func compile(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 || p[0].Type() != js.TypeString {
		return "invalid input"
	}

	input := p[0].String()

	tokens := lexer.Tokenize(input)
	parser := parser.New(tokens)
	ast := parser.Parse()

	result := generator.Compile(ast)
	return js.ValueOf(result)
}

func main() {
	js.Global().Set("compile", js.FuncOf(compile))
	select {}
}
