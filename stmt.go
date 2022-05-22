
package main


type StmtVisitorVoid interface {
    VisitExprStmt(expr *ExprStmt) 
	VisitPrintStmt(expr *PrintStmt) 
}

type StmtVisitor[R any] interface {
    VisitExprStmt(expr *ExprStmt) R
	VisitPrintStmt(expr *PrintStmt) R
}

type Stmt interface {
    AcceptString(visitor StmtVisitor[string]) string
	AcceptInterface(visitor StmtVisitor[interface{}]) interface{}
	Accept(visitor StmtVisitorVoid) 
}

type ExprStmt struct {
    Expression Expr
}

func (e *ExprStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitExprStmt(e)
}

func (e *ExprStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitExprStmt(e)
}

func (e *ExprStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitExprStmt(e)
}


type PrintStmt struct {
    Expression Expr
}

func (e *PrintStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitPrintStmt(e)
}

func (e *PrintStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitPrintStmt(e)
}

func (e *PrintStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitPrintStmt(e)
}



