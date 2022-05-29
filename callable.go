package main

import (
	"fmt"
	"time"
)

type Callable interface {
	Invoke(interpreter *Interpreter, args []interface{}) interface{}
	Arity() int
}

type ClockFunction struct{}

func (ClockFunction) Arity() int {
	return 0
}

func (ClockFunction) Invoke(i *Interpreter, args []interface{}) interface{} {
	return float64(time.Now().UnixMilli())
}

func (ClockFunction) String() string {
	return "<native func: clock>"
}

type Function struct {
	Definition *FunctionStmt
	Closure    *Environment
}

func (f *Function) Invoke(i *Interpreter, args []interface{}) (val interface{}) {
	defer func() {
		recover()
		val = GlobalEnv.Values["return"]
	}()
	env := NewEnvironment(f.Closure)
	for i, param := range f.Definition.Params {
		env.Define(string(param.Lexeme()), args[i])
	}

	i.executeBlock(f.Definition.Body, env)
	return
}

func (f *Function) Arity() int {
	return len(f.Definition.Params)
}

func (f *Function) String() string {
	return fmt.Sprintf("<func: %s>", string(f.Definition.Name.Lexeme()))
}
