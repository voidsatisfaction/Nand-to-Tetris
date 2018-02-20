package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	TypeA = "A"
	TypeC = "C"
	Label = "LABEL"
)

type Parser struct {
	filePath string
	result   []struct {
		commandType string
		fields      []string
	}
}

func New(filePath string) *Parser {
	return &Parser{filePath: filePath}
}

func (p *Parser) ParseFile() error {
	f, err := os.Open(p.filePath)
	s := bufio.NewScanner(f)
	if err != nil {
		return err
	}
	defer f.Close()

	for s.Scan() {
		line := s.Text()
		// remove strings after //
		if i := strings.Index(line, "//"); i >= 0 {
			line = line[:i]
		}
		// remove front, and back spaces
		line = strings.Trim(line, " ")
		// remove line starting with prefix // or empty line
		if strings.HasPrefix(line, "//") || line == "" {
			continue
		}

		// check label

		// check C command

		// check A command

		// throw error
		return errors.New("This is not valid command type " + fmt.Sprintf("%v\n", line))
	}

	return nil
}
