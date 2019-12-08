package asm

import (
	"fmt"
	"strings"
)

func decodeInstruction(code []byte, offset int) (string, int) {
	instruction := code[offset]
	opcode := Opcodes[instruction]

	operandList := make([]string, len(opcode.Operands))
	for i, operand := range opcode.Operands {
		operandList[i] = fmt.Sprintf("%s%x", operand.Type, code[offset+i+1])
	}

	return fmt.Sprintf("%s %s", opcode.Mnemonic, strings.Join(operandList, ", ")), offset+opcode.Length
}

func Disassemble(code []byte) {
	var line string
	i := 0
	for i < len(code) {
		line, i = decodeInstruction(code, i)
		fmt.Printf("%d: %s\n", i, line)
	}
}
