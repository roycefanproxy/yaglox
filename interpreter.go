package main

import (
	"fmt"

	"github.com/roycefanproxy/yaglox/constant"
)

type Interpreter struct {
	Env Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Env: Environment{
			Values: map[string]interface{}{},
		},
	}
}

func (i Interpreter) Interpret(statements []Stmt) {
	defer func() {
		recover()
	}()

	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i Interpreter) VisitLiteral(expr *Literal) interface{} {
	return expr.Value
}

func (i Interpreter) VisitGrouping(expr *Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i Interpreter) VisitUnary(expr *Unary) interface{} {
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type() {
	case constant.Minus:
		i.checkNumberOperand(expr.Operator, right)
		return -right.(float64)
	case constant.Bang:
		return !i.isTruthy(right)
	}

	return nil
}

func (i Interpreter) VisitBinary(expr *Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type() {
	case constant.BangEqual:
		return !i.isEqual(left, right)
	case constant.EqualEqual:
		return i.isEqual(left, right)
	case constant.Greater:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case constant.GreaterEqual:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case constant.Less:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case constant.LessEqual:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	case constant.Minus:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case constant.Plus:
		lNum, isLeftNum := left.(float64)
		rNum, isRightNum := right.(float64)
		if isLeftNum && isRightNum {
			return lNum + rNum
		}
		lStr, isLeftStr := left.(string)
		rStr, isRightStr := right.(string)
		if isLeftStr && isRightStr {
			return lStr + rStr
		}
		panic(i.error(expr.Operator, "Operands must be two numbers or two strings."))
	case constant.Slash:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case constant.Star:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	}

	return nil
}

func (i Interpreter) VisitVariable(expr *Variable) interface{} {
	return i.Env.Get(expr.Name)
}

func (i Interpreter) VisitExprStmt(stmt *ExprStmt) {
	i.evaluate(stmt.Expression)
}

func (i Interpreter) VisitPrintStmt(stmt *PrintStmt) {
	val := i.evaluate(stmt.Expression)
	fmt.Println(i.stringify(val))
}

func (i Interpreter) VisitVarDeclStmt(stmt *VarDeclStmt) {
	var val interface{}

	if stmt.Initializer != nil {
		val = i.evaluate(stmt.Initializer)
	}

	i.Env.Define(stmt.Name, val)
}

func (i Interpreter) evaluate(expr Expr) interface{} {
	return expr.AcceptInterface(i)
}

func (i Interpreter) execute(stmt Stmt) {
	stmt.Accept(i)
}

func (i Interpreter) isTruthy(val interface{}) bool {
	if val == nil {
		return false
	}

	switch v := val.(type) {
	case bool:
		return v
	case float64:
		return v != 0.0
	default:
		return true
	}
}

func (i Interpreter) isEqual(left, right interface{}) bool {
	return left == right
}

func (i Interpreter) checkNumberOperand(op Token, operand interface{}) {
	_, ok := operand.(float64)

	if !ok {
		panic(i.error(op, "Operand must be a number."))
	}
}

func (i Interpreter) checkNumberOperands(op Token, l, r interface{}) {
	_, lOk := l.(float64)
	_, rOk := r.(float64)
	if !lOk || !rOk {
		panic(i.error(op, "Operands must be numbers."))
	}
}

func (Interpreter) error(token Token, msg string) string {
	return EmitRuntimeError(token, msg)
}

func (Interpreter) stringify(val interface{}) string {
	if val == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", val)
}
