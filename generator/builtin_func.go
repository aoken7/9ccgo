package generator

var builtinFunc = map[string]string{
	"put": `	
put:
	push rbp
	mov rbp,rsp

	mov rax, 1
	mov rdi, [rbp+16]
	mov [rbp+16], rdi
	mov rdi, 1
	lea rsi, [rbp+16]
	mov rdx, 1
	syscall

	mov rsp,rbp
	pop rbp
	ret

`,
}
