package main

import "fmt"

type Token *TokenImpl

type TokenImpl struct {
	Type    TokenType
	Lexeme  []rune
	Literal interface{}
	Line    int
}

func (t TokenImpl) String() string {
	return fmt.Sprintf("<%s: %s>: %v", t.Type, string(t.Lexeme), t.Literal)
}

func NewToken(tokenType TokenType, lexeme []rune, literal interface{}, line int) Token {
	return &TokenImpl{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
