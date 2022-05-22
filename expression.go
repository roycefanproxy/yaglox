


package main

type Visitor[R any] interface {
    VisitBinaryExpr(expr *Binary) R
	VisitGroupingExpr(expr *Grouping) R
	VisitLiteralExpr(expr *Literal) R
	VisitUnaryExpr(expr *Unary) R
}

type Expr interface {
    AcceptString(visitor Visitor[string]) string
	AcceptInterface(visitor Visitor[interface{}]) interface{}
}

type Binary struct {
    Left Expr
	Operator Token
	Right Expr
}

func (e *Binary) AcceptString(visitor Visitor[string]) string {
    return visitor.VisitBinaryExpr(e)
}

func (e *Binary) AcceptInterface(visitor Visitor[interface{}]) interface{} {
    return visitor.VisitBinaryExpr(e)
}


type Grouping struct {
    Expression Expr
}

func (e *Grouping) AcceptString(visitor Visitor[string]) string {
    return visitor.VisitGroupingExpr(e)
}

func (e *Grouping) AcceptInterface(visitor Visitor[interface{}]) interface{} {
    return visitor.VisitGroupingExpr(e)
}


type Literal struct {
    Value interface{}
}

func (e *Literal) AcceptString(visitor Visitor[string]) string {
    return visitor.VisitLiteralExpr(e)
}

func (e *Literal) AcceptInterface(visitor Visitor[interface{}]) interface{} {
    return visitor.VisitLiteralExpr(e)
}


type Unary struct {
    Operator Token
	Right Expr
}

func (e *Unary) AcceptString(visitor Visitor[string]) string {
    return visitor.VisitUnaryExpr(e)
}

func (e *Unary) AcceptInterface(visitor Visitor[interface{}]) interface{} {
    return visitor.VisitUnaryExpr(e)
}



