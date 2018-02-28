package translator

import (
	"fmt"
	"strconv"
)

const (
	local    = "local"
	argument = "argument"
	this     = "this"
	that     = "that"
	constant = "constant"
	static   = "static"
	pointer  = "pointer"
	temp     = "temp"
)

func handleMemoryAccessCommand(tc translatorConfig, pc parsedCommand) (string, error) {
	ct := pc.CommandType()
	fields := pc.Fields()
	f1, f2 := fields[0], fields[1]

	switch ct {
	case push:
		return pushCommand(tc, f1, f2)
	case pop:
		return popCommand(tc, f1, f2)
	}
	return "", nil
}

func pushCommand(tc translatorConfig, f1, f2 string) (string, error) {
	switch f1 {
	case local:
		return localPush(f2), nil
	case argument:
		return argumentPush(f2), nil
	case this:
		return thisPush(f2), nil
	case that:
		return thatPush(f2), nil
	case constant:
		return constantPush(f2), nil
	case static:
		return staticPush(tc, f2), nil
	case pointer:
		return pointerPush(f2)
	case temp:
		return tempPush(f2)
	}
	return "", fmt.Errorf("this is not a proper memory area\n current value: %s", f1)
}

func popCommand(tc translatorConfig, f1, f2 string) (string, error) {
	switch f1 {
	case local:
		return localPop(f2), nil
	case argument:
		return argumentPop(f2), nil
	case this:
		return thisPop(f2), nil
	case that:
		return thatPop(f2), nil
	case static:
		return staticPop(tc, f2), nil
	case pointer:
		return pointerPop(f2)
	case temp:
		return tempPop(f2)
	}
	return "", fmt.Errorf("this is not a proper memory area\n current value: %s", f1)
}

// push commands
func localPush(f2 string) string {
	return "@LCL\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"A=D+A\n" +
		"D=M\n" +
		commonPush()
}

func argumentPush(f2 string) string {
	return "@ARG\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"A=D+A\n" +
		"D=M\n" +
		commonPush()
}

func thisPush(f2 string) string {
	return "@THIS\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"A=D+A\n" +
		"D=M\n" +
		commonPush()
}

func thatPush(f2 string) string {
	return "@THAT\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"A=D+A\n" +
		"D=M\n" +
		commonPush()
}

func staticPush(tc translatorConfig, f2 string) string {
	fileName := tc.BaseFileName()
	return "@" + fileName + "." + f2 + "\n" +
		"D=M\n" +
		commonPush()
}

func constantPush(f2 string) string {
	return "@" + f2 + "\n" +
		"D=A\n" +
		commonPush()
}

func tempPush(f2 string) (string, error) {
	tempAddr := 5
	numf2, err := strconv.Atoi(f2)
	if err != nil {
		return "", err
	}

	f2 = strconv.Itoa(tempAddr + numf2)
	if f2 >= "13" {
		return "", fmt.Errorf("temp address is overflowed! it should be more more than 5 and under 12\n current value is %d", tempAddr+numf2)
	}
	return "@" + f2 + "\n" +
		"D=M\n" +
		commonPush(), nil
}

func pointerPush(f2 string) (string, error) {
	var thisOrThat string
	switch f2 {
	case "0":
		thisOrThat = "THIS"
	case "1":
		thisOrThat = "THAT"
	default:
		return "", fmt.Errorf("pointer value is only 0 or 1\n current value: %s", f2)
	}
	return "@" + thisOrThat + "\n" +
		"D=M\n" +
		commonPush(), nil
}

// pop commands
func localPop(f2 string) string {
	return "@LCL\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"D=D+A\n" +
		commonPop()
}

func argumentPop(f2 string) string {
	return "@ARG\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"D=D+A\n" +
		commonPop()
}

func thisPop(f2 string) string {
	return "@THIS\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"D=D+A\n" +
		commonPop()
}

func thatPop(f2 string) string {
	return "@THAT\n" +
		"D=M\n" +
		"@" + f2 + "\n" +
		"D=D+A\n" +
		commonPop()
}

func staticPop(tc translatorConfig, f2 string) string {
	fileName := tc.BaseFileName()
	return "@" + fileName + "." + f2 + "\n" +
		"D=A\n" +
		commonPop()
}

func tempPop(f2 string) (string, error) {
	tempAddr := 5
	numf2, err := strconv.Atoi(f2)
	if err != nil {
		return "", err
	}

	f2 = strconv.Itoa(tempAddr + numf2)
	if f2 >= "13" {
		return "", fmt.Errorf("temp address is overflowed! it should be more more than 5 and under 12\n current value is %d", tempAddr+numf2)
	}
	return "@" + f2 + "\n" +
		"D=A\n" +
		commonPop(), nil
}

func pointerPop(f2 string) (string, error) {
	var thisOrThat string
	switch f2 {
	case "0":
		thisOrThat = "THIS"
	case "1":
		thisOrThat = "THAT"
	default:
		return "", fmt.Errorf("pointer value is only 0 or 1\n current value: %s", f2)
	}
	return "@" + thisOrThat + "\n" +
		"D=A\n" +
		commonPop(), nil
}

// common logics
func commonPush() string {
	return "@SP\n" +
		"A=M\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M+1\n"
}

func commonPop() string {
	return "@R15\n" +
		"M=D\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@R15\n" +
		"A=M\n" +
		"M=D\n"
}
