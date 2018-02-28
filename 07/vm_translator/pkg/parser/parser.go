package parser

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Push = "push"
	Pop  = "pop"

	Add = "add"
	Sub = "sub"
)

type parsedCommand struct {
	commandType string
	fields      []string
}

func newParsedCommand(commandType string, fields ...string) *parsedCommand {
	return &parsedCommand{
		commandType, fields,
	}
}

func (pc *parsedCommand) CommandType() string {
	return pc.commandType
}

func (pc *parsedCommand) Fields() []string {
	return pc.fields
}

type Parser struct {
	parsedCommands []*parsedCommand
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParsedCommands() []*parsedCommand {
	return p.parsedCommands
}

func (p *Parser) Parse(line string) error {
	line = removeWhiteSpace(line)
	if line == "" {
		return nil
	}

	cmd := strings.Split(line, " ")
	if len(cmd) < 1 {
		return errors.New(fmt.Sprintf("Parse: command should have more than one field\n line: %s", line))
	}

	var commandType string
	var fields []string

	commandType = cmd[0]
	if len(cmd) > 1 {
		fields = cmd[1:]
	}
	p.addCommandToParser(commandType, fields...)
	return nil
}

func removeWhiteSpace(rawString string) string {
	if strings.HasPrefix(rawString, "//") || rawString == "" {
		return ""
	}
	// remove strings after //
	if i := strings.Index(rawString, "//"); i >= 0 {
		rawString = rawString[:i]
	}
	// remove front, and back spaces
	rawString = strings.Trim(rawString, " ")

	return rawString
}

func (p *Parser) addCommandToParser(commandType string, fields ...string) {
	pc := newParsedCommand(commandType, fields...)
	p.parsedCommands = append(p.parsedCommands, pc)
}
