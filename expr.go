
package main


type ExprVisitorVoid interface {
    VisitBinary(expr *Binary) 
	VisitGrouping(expr *Grouping) 
	VisitLiteral(expr *Literal) 
	VisitUnary(expr *Unary) 
}

type ExprVisitor[R any] interface {
    VisitBinary(expr *Binary) R
	VisitGrouping(expr *Grouping) R
	VisitLiteral(expr *Literal) R
	VisitUnary(expr *Unary) R
}

type Expr interface {
    AcceptString(visitor ExprVisitor[string]) string
	AcceptInterface(visitor ExprVisitor[interface{}]) interface{}
	Accept(visitor ExprVisitorVoid) 
}

type Binary struct {
    Left Expr
	Operator Token
	Right Expr
}

func (e *Binary) AcceptString(visitor ExprVisitor[string]) string {
    return visitor.VisitBinary(e)
}

func (e *Binary) AcceptInterface(visitor ExprVisitor[interface{}]) interface{} {
    return visitor.VisitBinary(e)
}

func (e *Binary) Accept(visitor ExprVisitorVoid)  {
    visitor.VisitBinary(e)
}


type Grouping struct {
    Expression Expr
}

func (e *Grouping) AcceptString(visitor ExprVisitor[string]) string {
    return visitor.VisitGrouping(e)
}

func (e *Grouping) AcceptInterface(visitor ExprVisitor[interface{}]) interface{} {
    return visitor.VisitGrouping(e)
}

func (e *Grouping) Accept(visitor ExprVisitorVoid)  {
    visitor.VisitGrouping(e)
}


type Literal struct {
    Value interface{}
}

func (e *Literal) AcceptString(visitor ExprVisitor[string]) string {
    return visitor.VisitLiteral(e)
}

func (e *Literal) AcceptInterface(visitor ExprVisitor[interface{}]) interface{} {
    return visitor.VisitLiteral(e)
}

func (e *Literal) Accept(visitor ExprVisitorVoid)  {
    visitor.VisitLiteral(e)
}


type Unary struct {
    Operator Token
	Right Expr
}

func (e *Unary) AcceptString(visitor ExprVisitor[string]) string {
    return visitor.VisitUnary(e)
}

func (e *Unary) AcceptInterface(visitor ExprVisitor[interface{}]) interface{} {
    return visitor.VisitUnary(e)
}

func (e *Unary) Accept(visitor ExprVisitorVoid)  {
    visitor.VisitUnary(e)
}



