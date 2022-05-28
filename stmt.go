
package main


type StmtVisitorVoid interface {
    VisitExprStmt(expr *ExprStmt) 
	VisitPrintStmt(expr *PrintStmt) 
	VisitVarDeclStmt(expr *VarDeclStmt) 
}

type StmtVisitor[R any] interface {
    VisitExprStmt(expr *ExprStmt) R
	VisitPrintStmt(expr *PrintStmt) R
	VisitVarDeclStmt(expr *VarDeclStmt) R
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


type VarDeclStmt struct {
    Name Token
	Initializer Expr
}

func (e *VarDeclStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitVarDeclStmt(e)
}

func (e *VarDeclStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitVarDeclStmt(e)
}

func (e *VarDeclStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitVarDeclStmt(e)
}



