package main

import (
	"github.com/roycefanproxy/yaglox/constant"
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

func (p *Parser) Parse() []Stmt {
	defer func() {
		recover()
	}()

	stmts := make([]Stmt, 0)

	for !p.isAtEnd() {
		stmts = append(stmts, p.declaration())
	}

	return stmts
}

/*
func (p *Parser) Parse() (expr Expr) {
	defer func() {
		recover()
	}()

	expr = p.expression()
	return
}
*/

func (p *Parser) declaration() Stmt {
	defer func() {
		if err := recover(); err != nil {
			p.Synchronize()
		}
	}()

	if p.match(constant.Var) {
		return p.varDeclaration()
	}

	return p.statement()
}

func (p *Parser) varDeclaration() Stmt {
	name := p.consume(constant.Identifier, "Expect variable name.")

	var initializer Expr
	if p.match(constant.Equal) {
		initializer = p.expression()
	}

	p.consume(constant.Semicolon, "Expect ';' after variable declaration.")

	return &VarDeclStmt{
		Name:        name,
		Initializer: initializer,
	}
}

func (p *Parser) statement() Stmt {
	if p.match(constant.Print) {
		return p.printStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) printStatement() Stmt {
	expr := p.expression()
	p.consume(constant.Semicolon, "Expect ';' after value.")
	return &PrintStmt{
		Expression: expr,
	}
}

func (p *Parser) expressionStatement() Stmt {
	expr := p.expression()
	p.consume(constant.Semicolon, "Expect ';' after value.")
	return &ExprStmt{
		Expression: expr,
	}
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	return p.binaryMatcher((*Parser).comparison, constant.BangEqual, constant.EqualEqual)
}

func (p *Parser) comparison() Expr {
	return p.binaryMatcher((*Parser).term, constant.Less, constant.LessEqual, constant.GreaterEqual, constant.Greater)
}

func (p *Parser) term() Expr {
	return p.binaryMatcher((*Parser).factor, constant.Minus, constant.Plus)
}

func (p *Parser) factor() Expr {
	return p.binaryMatcher((*Parser).unary, constant.Slash, constant.Star)
}

func (p *Parser) unary() Expr {
	if p.match(constant.Bang, constant.Minus) {
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
	case constant.False:
		expr = &Literal{Value: false}
	case constant.True:
		expr = &Literal{Value: true}
	case constant.Nil:
		expr = &Literal{Value: nil}
	case constant.Number, constant.String:
		expr = &Literal{Value: p.peek().Literal()}
	case constant.Identifier:
		expr = &Variable{Name: p.peek()}
	case constant.LeftParen:
		p.advance()
		expr = p.expression()
		p.consume(constant.RightParen, "Expect ')' after expression.")
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

func (p *Parser) consume(tokenType constant.TokenType, msg string) Token {
	if p.check(tokenType) {
		return p.advance()
	}

	panic(p.error(p.peek(), msg))
}

func (p *Parser) match(tokenTypes ...constant.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType constant.TokenType) bool {
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
	return p.peek().Type() == constant.EOF
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
		if p.previous().Type() == constant.Semicolon {
			return
		}

		switch p.peek().Type() {
		case constant.Class, constant.Func, constant.Var, constant.For,
			constant.If, constant.While, constant.Print, constant.Return:
			return
		}

		p.advance()
	}
}

func (p *Parser) binaryMatcher(operand func(p *Parser) Expr, matches ...constant.TokenType) Expr {
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
