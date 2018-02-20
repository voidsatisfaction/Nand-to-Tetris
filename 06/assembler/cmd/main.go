package main

import (
	"fmt"

	"../pkg/parser"
)

func main() {
	p := parser.New(`../../add/Add.asm`)
	if err := p.ParseFile(); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", p)
}
