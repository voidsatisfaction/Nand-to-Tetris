package translator

import (
	"errors"
	"fmt"
)

var (
	eqIndex = -1
	gtIndex = -1
	ltIndex = -1
)

func handleArithmeticLogicCommand(pc parsedCommand) (string, error) {
	ct := pc.CommandType()
	switch ct {
	case add:
		return addCommand(), nil
	case sub:
		return subCommand(), nil
	case neg:
		return negCommand(), nil
	case and:
		return andCommand(), nil
	case or:
		return orCommand(), nil
	case not:
		return notCommand(), nil
	case eq:
		return eqCommand(), nil
	case gt:
		return gtCommand(), nil
	case lt:
		return ltCommand(), nil
	}
	return "", errors.New(fmt.Sprintf("This is not valid command\ncommand: %v", ct))
}

func addCommand() string {
	return `@SP
AM=M-1
D=M
A=A-1
M=D+M
`
}

func subCommand() string {
	return `@SP
AM=M-1
D=M
A=A-1
M=M-D
`
}

func negCommand() string {
	return `@SP
A=M-1
M=-M
`
}

func eqCommand() string {
	eqIndex++
	eqTrueLabel := fmt.Sprintf("EQ.true.%d", eqIndex)
	eqFalseLabel := fmt.Sprintf("EQ.false.%d", eqIndex)
	eqAfterLabel := fmt.Sprintf("EQ.after.%d", eqIndex)
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@" + eqTrueLabel + "\n" +
		"D;JEQ\n" +
		"@" + eqFalseLabel + "\n" +
		"0;JMP\n" +
		"(" + eqTrueLabel + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"@" + eqAfterLabel + "\n" +
		"0;JMP\n" +
		"(" + eqFalseLabel + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"(" + eqAfterLabel + ")\n"
}

func gtCommand() string {
	gtIndex++
	gtTrueLabel := fmt.Sprintf("GT.true.%d", gtIndex)
	gtFalseLabel := fmt.Sprintf("GT.false.%d", gtIndex)
	gtAfterLabel := fmt.Sprintf("GT.after.%d", gtIndex)
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@" + gtTrueLabel + "\n" +
		"D;JGT\n" +
		"@" + gtFalseLabel + "\n" +
		"0;JMP\n" +
		"(" + gtTrueLabel + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"@" + gtAfterLabel + "\n" +
		"0;JMP\n" +
		"(" + gtFalseLabel + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"(" + gtAfterLabel + ")\n"
}

func ltCommand() string {
	ltIndex++
	ltTrueLabel := fmt.Sprintf("LT.true.%d", ltIndex)
	ltFalseLabel := fmt.Sprintf("LT.false.%d", ltIndex)
	ltAfterLabel := fmt.Sprintf("LT.after.%d", ltIndex)
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@" + ltTrueLabel + "\n" +
		"D;JLT\n" +
		"@" + ltFalseLabel + "\n" +
		"0;JMP\n" +
		"(" + ltTrueLabel + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"@" + ltAfterLabel + "\n" +
		"0;JMP\n" +
		"(" + ltFalseLabel + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"(" + ltAfterLabel + ")\n"
}

func andCommand() string {
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D&M\n"
}

func orCommand() string {
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D|M\n"
}

func notCommand() string {
	return "@SP\n" +
		"A=M-1\n" +
		"M=!M\n"
}
