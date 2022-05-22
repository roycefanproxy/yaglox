package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ASTPrinter struct{}

func (p ASTPrinter) Print(expr Expr) string {
	return expr.AcceptString(p)
}

func (p ASTPrinter) VisitBinary(expr *Binary) string {
	return p.parenthesize(expr.Operator.Lexeme(), expr.Left, expr.Right)
}

func (p ASTPrinter) VisitGrouping(expr *Grouping) string {
	return p.parenthesize([]rune("group"), expr.Expression)
}

func (p ASTPrinter) VisitLiteral(expr *Literal) string {
	if expr.Value == nil {
		return "nil"
	}

	switch val := expr.Value.(type) {
	case float64:
		return fmt.Sprintf("%f", val)
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("unknown type: %v", val)
	}
}

func (p ASTPrinter) VisitUnary(expr *Unary) string {
	return p.parenthesize(expr.Operator.Lexeme(), expr.Right)
}

func (p ASTPrinter) parenthesize(name []rune, expressions ...Expr) string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(string(name))
	for _, expression := range expressions {
		str := fmt.Sprintf(" %s", expression.AcceptString(p))
		builder.WriteString(str)
	}
	builder.WriteString(")")

	return builder.String()
}
