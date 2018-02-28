package translator

import (
	"fmt"
)

const (
	add = "add"
	sub = "sub"
	neg = "neg"
	eq  = "eq"
	gt  = "gt"
	lt  = "lt"
	and = "and"
	or  = "or"
	not = "not"

	pop  = "pop"
	push = "push"
)

type translatorConfig interface {
	BaseFileName() string
}

type parsedCommand interface {
	CommandType() string
	Fields() []string
}

func Translate(tc translatorConfig, pc parsedCommand) (string, error) {
	ct := pc.CommandType()
	if checkArithmeticLogicCommand(ct) {
		return handleArithmeticLogicCommand(pc)
	}

	if checkMemoryAccessCommand(ct) {
		return handleMemoryAccessCommand(tc, pc)
	}

	return "", fmt.Errorf("translate: this command type is not valid\n command: %+v", pc.CommandType())
}

func checkArithmeticLogicCommand(commandType string) bool {
	switch commandType {
	case add, sub, neg, eq, gt, lt, and, or, not:
		return true
	}
	return false
}

func checkMemoryAccessCommand(commandType string) bool {
	switch commandType {
	case pop, push:
		return true
	}
	return false
}
