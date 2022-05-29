package main

import "fmt"

type Environment struct {
	Values   map[string]interface{}
	OuterEnv *Environment
}

func NewEnvironment(outerEnv *Environment) *Environment {
	return &Environment{
		Values:   map[string]interface{}{},
		OuterEnv: outerEnv,
	}
}

func (env *Environment) Define(name Token, value interface{}) {
	env.Values[string(name.Lexeme())] = value
}

func (env *Environment) Get(name Token) interface{} {
	val, hasVal := env.Values[string(name.Lexeme())]
	if hasVal {
		return val
	}

	if env.OuterEnv != nil {
		return env.OuterEnv.Get(name)
	}

	errMsg := fmt.Sprintf("Undefined variable '%s'.", string(name.Lexeme()))
	panic(EmitRuntimeError(name, errMsg))
}

func (env *Environment) Assign(name Token, value interface{}) {
	nameStr := string(name.Lexeme())
	if _, ok := env.Values[nameStr]; ok {
		env.Values[nameStr] = value
		return
	}

	if env.OuterEnv != nil {
		env.OuterEnv.Assign(name, value)
		return
	}

	msg := fmt.Sprintf("Undefined variable '%s'.", nameStr)
	panic(EmitRuntimeError(name, msg))
}
