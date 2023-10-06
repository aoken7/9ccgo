# Linuxで動作するx86-64のアセンブリコードをエンターを押すごとに読み込み実行する
# また、スタックの状態とレジスタの状態をCUIで整形して表示する

stack = ['x' for i in range(30)]
registers = {
    'rax': 0,
    'rbx': 0,
    'rcx': 0,
    'rdx': 0,
    'rsi': 0,
    'rdi': 0,
    'rbp': 0,
    'rsp': 0,
    'rip': 0,
    'r8': 0,
    'r9': 0,
    'r10': 0,
    'r11': 0,
    'r12': 0,
    'r13': 0,
    'r14': 0,
    'r15': 0,
    'rflags': 0,
}

def print_stack():
    print('stack: ', end='')
    for i in stack:
        print(i, end=' ')
    
def print_rax_rdi_rbp_rsp():
    print('rax: {:<5}  rdi: {:<5}  rbp: {:<5}  rsp: {:<5}'.format(registers['rax'], registers['rdi'], registers['rbp'], registers['rsp']))


def emurate(line: list):
    if len(line) == 1:
        return
    
    opecode = line[0].strip()
    operand1 = line[1].strip().replace(',', '')

    if opecode == 'push':
        if operand1 in registers:
            registers['rsp'] -= 8
            stack[-registers['rsp']//8] = registers[operand1]
        else:
            registers['rsp'] -= 8
            stack[-registers['rsp']//8] = int(operand1)
    elif opecode == 'pop':    
        registers[operand1] = stack[-registers['rsp']//8]
        registers['rsp'] += 8

    elif opecode == 'mov':
        operand2 = line[2]
        if operand1[0] == '[':
            operand1 = operand1[1:-1]
            stack[-registers[operand1]//8] = registers[operand2]
        elif operand2[0] == '[':
            operand2 = line[2][1:-1]
            registers[operand1] = stack[-registers[operand2]//8]
        else:
            registers[operand1] = registers[operand2]

    elif opecode == 'add':
        operand2 = line[2]
        if operand2 in registers:
            registers[operand1] += registers[operand2]
        else:
            registers[operand1] += int(operand2)
    elif opecode == 'sub':
        operand2 = line[2]
        if operand2 in registers:
            registers[operand1] -= registers[operand2]
        else:
            registers[operand1] -= int(operand2)
    elif opecode == 'ret':
        return

def read_line(code: list):
    for line in code:
        if line == '':
            continue
        emurate(line.split(' '))
        print(line[1:])
        print_stack()
        print()
        print_rax_rdi_rbp_rsp()
        input()
        

        

def main():
    with open('tmp.s', 'r') as f:
        code = f.read()
    read_line(code.split('\n')[3:])
    
if __name__ == '__main__':
    main()