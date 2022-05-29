package main

import (
	"fmt"

	"github.com/roycefanproxy/yaglox/constant"
)

type Interpreter struct {
	Env *Environment
}

func NewInterpreter() *Interpreter {

	return &Interpreter{
		Env: GlobalEnv,
	}
}

func (i *Interpreter) Interpret(statements []Stmt) {
	defer func() {
		recover()
	}()

	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) VisitLiteral(expr *Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGrouping(expr *Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnary(expr *Unary) interface{} {
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

func (i *Interpreter) VisitBinary(expr *Binary) interface{} {
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

func (i *Interpreter) VisitCall(expr *Call) interface{} {
	callee := i.evaluate(expr.Callee)

	args := []interface{}{}
	for _, arg := range expr.Arguments {
		args = append(args, i.evaluate(arg))
	}

	callable, ok := callee.(Callable)
	if !ok {
		EmitRuntimeError(expr.Operator, "Can only call functions and classes.")
	}

	if argsLen, arity := len(args), callable.Arity(); argsLen != arity {
		msg := fmt.Sprintf("Expected %v arguments but got %v.", arity, argsLen)
		panic(EmitRuntimeError(expr.Operator, msg))
	}
	return callable.Invoke(i, args)
}

func (i *Interpreter) VisitLogical(expr *Logical) interface{} {
	left := i.evaluate(expr.Left)

	if isLeftTruthy := i.isTruthy(left); expr.Operator.Type() == constant.Or {
		if isLeftTruthy {
			return left
		}
	} else {
		if !isLeftTruthy {
			return left
		}
	}

	return i.evaluate(expr.Right)
}

func (i *Interpreter) VisitVariable(expr *Variable) interface{} {
	return i.Env.Get(expr.Name)
}

func (i *Interpreter) VisitAssign(expr *Assign) interface{} {
	value := i.evaluate(expr.Value)
	i.Env.Assign(expr.Name, value)
	return value
}

func (i *Interpreter) VisitExprStmt(stmt *ExprStmt) {
	i.evaluate(stmt.Expression)
}

func (i *Interpreter) VisitFunctionStmt(stmt *FunctionStmt) {
	function := &Function{
		Definition: stmt,
		Closure:    i.Env,
	}
	i.Env.Define(string(stmt.Name.Lexeme()), function)
}

func (i *Interpreter) VisitReturnStmt(stmt *ReturnStmt) {
	var val interface{}
	if stmt.Value != nil {
		val = i.evaluate(stmt.Value)
		GlobalEnv.Values["return"] = val
	}
	panic("return")
}

func (i *Interpreter) VisitPrintStmt(stmt *PrintStmt) {
	val := i.evaluate(stmt.Expression)
	fmt.Println(i.stringify(val))
}

func (i *Interpreter) VisitVarDeclStmt(stmt *VarDeclStmt) {
	var val interface{}

	if stmt.Initializer != nil {
		val = i.evaluate(stmt.Initializer)
	}

	i.Env.Define(string(stmt.Name.Lexeme()), val)
}

func (i *Interpreter) VisitBlockStmt(stmt *BlockStmt) {
	env := NewEnvironment(i.Env)
	i.executeBlock(stmt.Statements, env)
}

func (i *Interpreter) VisitIfStmt(stmt *IfStmt) {
	if i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.Then)
	} else if stmt.Else != nil {
		i.execute(stmt.Else)
	}
}

func (i *Interpreter) VisitWhileStmt(stmt *WhileStmt) {
	for i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.Statement)
	}
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return expr.AcceptInterface(i)
}

func (i *Interpreter) execute(stmt Stmt) {
	stmt.Accept(i)
}

func (i *Interpreter) executeBlock(statements []Stmt, env *Environment) {
	prevEnv := i.Env
	defer func() {
		recover()
		i.Env = prevEnv
	}()

	i.Env = env
	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) isTruthy(val interface{}) bool {
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

func (i *Interpreter) isEqual(left, right interface{}) bool {
	return left == right
}

func (i *Interpreter) checkNumberOperand(op Token, operand interface{}) {
	_, ok := operand.(float64)

	if !ok {
		panic(i.error(op, "Operand must be a number."))
	}
}

func (i *Interpreter) checkNumberOperands(op Token, l, r interface{}) {
	_, lOk := l.(float64)
	_, rOk := r.(float64)
	if !lOk || !rOk {
		panic(i.error(op, "Operands must be numbers."))
	}
}

func (*Interpreter) error(token Token, msg string) string {
	return EmitRuntimeError(token, msg)
}

func (*Interpreter) stringify(val interface{}) string {
	if val == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", val)
}
