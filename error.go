package main

import (
	"fmt"
	"os"

	tokentype "github.com/roycefanproxy/yaglox/constant"
)

var (
	HasError        = false
	HasRuntimeError = false
)

func EmitErrorLog(line int, msg string) {
	report(line, "", msg)
}

func report(line int, location, msg string) string {
	errMsg := fmt.Sprintf("[line %d] Error %s: %s", line, location, msg)
	fmt.Fprintln(os.Stderr, errMsg)
	HasError = true
	return errMsg
}

func EmitParseError(token Token, msg string) string {
	loc := " at end"
	if token.Type() != tokentype.EOF {
		loc = fmt.Sprintf("at ' %s'", string(token.Lexeme()))
	}

	return report(token.Line(), loc, msg)
}

func EmitRuntimeError(token Token, msg string) string {
	str := fmt.Sprintf("%s\n[line %v]", msg, token.Line())
	HasRuntimeError = true
	fmt.Println(str)

	return str
}
