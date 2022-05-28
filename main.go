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

// func main() {
// 	expression := &Binary{
// 		Left: &Unary{
// 			Operator: NewToken(Minus, []rune("-"), nil, 1),
// 			Right:    &Literal{Value: 123},
// 		},
// 		Operator: NewToken(Star, []rune("*"), nil, 1),
// 		Right: &Grouping{
// 			Expression: &Literal{Value: 45.67},
// 		},
// 	}
// 	fmt.Println(ASTPrinter{}.Print(expression))
// }

func execute(source string) {
	tokenizer := NewTokenizer(source)
	tokens := tokenizer.Parse()
	parser := NewParser(tokens)
	statements := parser.Parse()

	if HasError {
		return
	}

	/*
		fmt.Println("tokens: ")
		for _, t := range tokens {
			fmt.Printf("%s ", t)
		}
		fmt.Println()
		fmt.Println(ASTPrinter{}.Print(expression))
	*/

	NewInterpreter().Interpret(statements)
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
	if HasRuntimeError {
		os.Exit(70)
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
