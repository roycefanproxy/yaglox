package main

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
	return p.binaryMatcher((*Parser).comparison, BangEqual, EqualEqual)
}

func (p *Parser) comparison() Expr {
	return p.binaryMatcher((*Parser).term, Less, LessEqual, GreaterEqual, Greater)
}

func (p *Parser) term() Expr {
	return p.binaryMatcher((*Parser).factor, Minus, Plus)
}

func (p *Parser) factor() Expr {
	return p.binaryMatcher((*Parser).unary, Slash, Star)
}

func (p *Parser) unary() Expr {
	if p.match(Bang, Minus) {
		operator := p.previous()
		right := p.unary()
		return &Unary{
			Operator: operator,
			Right:    right,
		}
	}

	return p.primary()
}

func (p *Parser) primary() Expr {
	switch p.peek().Type {
	case False:
		return &Literal{Value: false}
	case True:
		return &Literal{Value: true}
	case Nil:
		return &Literal{Value: nil}
	case Number, String:
		return &Literal{p.previous().Literal}
	case LeftParen:
		expr := p.expression()
		p.consume(RightParen, "Expect ')' after expression.")
		return &Grouping{Expression: expr}
	}

	panic(p.error(p.peek(), "Expect expression."))
}

func (p *Parser) consume(tokenType TokenType, msg string) Token {
	if p.check(tokenType) {
		return p.advance()
	}

	panic(p.error(p.peek(), msg))
}

func (p *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) error(token Token, msg string) string {
	return EmitParseError(token, msg)
}

func (p *Parser) binaryMatcher(operand func(p *Parser) Expr, matches ...TokenType) Expr {
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
