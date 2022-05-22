package main

import (
	tokentype "github.com/roycefanproxy/yaglox/constant"
)

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() (expr Expr) {
	defer func() {
		recover()
	}()

	expr = p.expression()
	return
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	return p.binaryMatcher((*Parser).comparison, tokentype.BangEqual, tokentype.EqualEqual)
}

func (p *Parser) comparison() Expr {
	return p.binaryMatcher((*Parser).term, tokentype.Less, tokentype.LessEqual, tokentype.GreaterEqual, tokentype.Greater)
}

func (p *Parser) term() Expr {
	return p.binaryMatcher((*Parser).factor, tokentype.Minus, tokentype.Plus)
}

func (p *Parser) factor() Expr {
	return p.binaryMatcher((*Parser).unary, tokentype.Slash, tokentype.Star)
}

func (p *Parser) unary() Expr {
	if p.match(tokentype.Bang, tokentype.Minus) {
		operator := p.previous()
		right := p.unary()
		return &Unary{
			Operator: operator,
			Right:    right,
		}
	}

	return p.primary()
}

func (p *Parser) primary() (expr Expr) {
	if p.isAtEnd() {
		panic(p.error(p.peek(), "Expect expression."))
	}

	switch p.peek().Type() {
	case tokentype.False:
		expr = &Literal{Value: false}
	case tokentype.True:
		expr = &Literal{Value: true}
	case tokentype.Nil:
		expr = &Literal{Value: nil}
	case tokentype.Number, tokentype.String:
		expr = &Literal{p.peek().Literal()}
	case tokentype.LeftParen:
		p.advance()
		expr := p.expression()
		p.consume(tokentype.RightParen, "Expect ')' after expression.")
		expr = &Grouping{Expression: expr}
		goto post_advance
	}
	p.advance()

post_advance:
	if expr == nil {
		panic(p.error(p.peek(), "Expect expression."))
	}

	return
}

/*
func (p *Parser) primary() (expr Expr) {
	if p.match(False) {
		return &Literal{Value: false}
	}

	if p.match(True) {
		return &Literal{Value: true}
	}

	if p.match(Number, String) {
		return &Literal{Value: p.previous().Literal()}
	}

	if p.match(LeftParen) {
		expr := p.expression()
		p.consume(RightParen, "Expect ')' after expression.")
		return &Grouping{Expression: expr}
	}

	panic(p.error(p.peek(), "Expect expression."))
}
*/

func (p *Parser) consume(tokenType tokentype.TokenType, msg string) Token {
	if p.check(tokenType) {
		return p.advance()
	}

	panic(p.error(p.peek(), msg))
}

func (p *Parser) match(tokenTypes ...tokentype.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType tokentype.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type() == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type() == tokentype.EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (*Parser) error(token Token, msg string) string {
	return EmitParseError(token, msg)
}

func (p *Parser) Synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type() == tokentype.Semicolon {
			return
		}

		switch p.peek().Type() {
		case tokentype.Class, tokentype.Func, tokentype.Var, tokentype.For,
			tokentype.If, tokentype.While, tokentype.Print, tokentype.Return:
			return
		}

		p.advance()
	}
}

func (p *Parser) binaryMatcher(operand func(p *Parser) Expr, matches ...tokentype.TokenType) Expr {
	expr := operand(p)

	for p.match(matches...) {
		operator := p.previous()
		right := operand(p)
		bin := &Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
		expr = bin
	}

	return expr
}
