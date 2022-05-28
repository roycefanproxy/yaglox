package main

import "fmt"

type Environment struct {
	Values map[string]interface{}
}

func (env *Environment) Define(name Token, value interface{}) {
	env.Values[string(name.Lexeme())] = value
}

func (env *Environment) Get(name Token) interface{} {
	val, hasVal := env.Values[string(name.Lexeme())]
	if !hasVal {
		errMsg := fmt.Sprintf("Undefined variable '%s'.", string(name.Lexeme()))
		panic(EmitRuntimeError(name, errMsg))
	}

	return val
}
