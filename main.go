package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 2:
		executeFile(os.Args[1])
	case 1:
		executePrompt()
	default:
		fmt.Print("Usage: lox [script]")
		os.Exit(64)
	}

}

func execute(source string) {
	tokenizer := NewTokenizer(source)
	tokens := tokenizer.Parse()
	parser := NewParser(tokens)
	parser.Parse()

	if HasError {
		return
	}

}

func executeFile(filePath string) {
	bin, err := os.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	execute(string(bin))

	if HasError {
		os.Exit(65)
	}
}

func executePrompt() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("->")
		if !reader.Scan() {
			break
		}

		execute(reader.Text())
		HasError = false
	}

}
