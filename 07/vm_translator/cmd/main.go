package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"../pkg/parser"
	"../pkg/translator"
)

const (
	Ext = ".asm"
)

type hackVMTranslatorConfig struct {
	baseFileName string
}

func newHackVMTranslatorConfig(filePath string) *hackVMTranslatorConfig {
	fn := filepath.Base(filePath)
	return &hackVMTranslatorConfig{
		baseFileName: fn,
	}
}

func (tc *hackVMTranslatorConfig) BaseFileName() string {
	return tc.baseFileName
}

type hackVMTranslator struct {
	*parser.Parser
}

func newHackVMTranslator() *hackVMTranslator {
	return &hackVMTranslator{
		Parser: parser.New(),
	}
}

func main() {
	filePath := os.Args[1]
	fc := newHackVMTranslatorConfig(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	hvmt := newHackVMTranslator()
	for s.Scan() {
		line := s.Text()

		err := hvmt.Parser.Parse(line)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}

	filePathUntilExt := strings.LastIndex(filePath, ".")
	outFile, err := os.Create(filePath[:filePathUntilExt] + Ext)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	w := bufio.NewWriter(outFile)
	for _, parsedCommand := range hvmt.Parser.ParsedCommands() {
		asmCommand, err := translator.Translate(fc, parsedCommand)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		w.WriteString(asmCommand)
		fmt.Println(asmCommand)
	}
	w.Flush()
}
