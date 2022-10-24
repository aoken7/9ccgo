package main

import (
	"fmt"
	"os"
	"strconv"
)

func compile(s string) string {
	num, _ := strconv.Atoi(s)
	t := fmt.Sprintf(".intel_syntax noprefix\n")
	t += fmt.Sprintf(".global main\n")
	t += fmt.Sprintf("main:\n")
	t += fmt.Sprintf("\tmov rax, %d\n", num)
	t += fmt.Sprintf("\tret\n")
	return t
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough args.")
		return
	}

	t := compile(os.Args[1])
	fmt.Println(t)
}
