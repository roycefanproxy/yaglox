package main

import (
	"fmt"

	tokentype "github.com/roycefanproxy/yaglox/constant"
)

type Token interface {
	fmt.Stringer
	Type() tokentype.TokenType
	Lexeme() []rune
	Literal() interface{}
	Line() int
}

type TokenImpl struct {
	tokenType tokentype.TokenType
	lexeme    []rune
	literal   interface{}
	line      int
}

func (t *TokenImpl) Type() tokentype.TokenType {
	return t.tokenType
}

func (t *TokenImpl) Lexeme() []rune {
	return t.lexeme
}

func (t *TokenImpl) Literal() interface{} {
	return t.literal
}

func (t *TokenImpl) Line() int {
	return t.line
}

func (t *TokenImpl) String() (str string) {
	strLexeme := string(t.Lexeme())
	if strLexeme == "" {
		str = fmt.Sprintf("<%s>", t.Type())
	} else {
		str = fmt.Sprintf("<%s: %s>", t.Type(), string(t.Lexeme()))
	}
	if t.Literal() != nil {
		str = fmt.Sprintf("%s: %v", str, t.Literal())
	}

	return str
}

func NewToken(tokenType tokentype.TokenType, lexeme []rune, literal interface{}, line int) Token {
	return &TokenImpl{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}
