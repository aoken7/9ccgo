package generator

var builtinFunc = map[string]string{
	"put": `	mov rax, 1
	
	mov rdi, [rbp+16]
	add rdi, 48
	mov [rbp+16], rdi
	
	mov rdi, 1
	lea rsi, [rbp+16]
	mov rdx, 1
	syscall
`,
}
