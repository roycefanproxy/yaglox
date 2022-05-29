
package main


type StmtVisitorVoid interface {
    VisitExprStmt(expr *ExprStmt) 
	VisitFunctionStmt(expr *FunctionStmt) 
	VisitIfStmt(expr *IfStmt) 
	VisitWhileStmt(expr *WhileStmt) 
	VisitVarDeclStmt(expr *VarDeclStmt) 
	VisitBlockStmt(expr *BlockStmt) 
	VisitReturnStmt(expr *ReturnStmt) 
	VisitPrintStmt(expr *PrintStmt) 
}

type StmtVisitor[R any] interface {
    VisitExprStmt(expr *ExprStmt) R
	VisitFunctionStmt(expr *FunctionStmt) R
	VisitIfStmt(expr *IfStmt) R
	VisitWhileStmt(expr *WhileStmt) R
	VisitVarDeclStmt(expr *VarDeclStmt) R
	VisitBlockStmt(expr *BlockStmt) R
	VisitReturnStmt(expr *ReturnStmt) R
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


type FunctionStmt struct {
    Name Token
	Params []Token
	Body []Stmt
}

func (e *FunctionStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitFunctionStmt(e)
}

func (e *FunctionStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitFunctionStmt(e)
}

func (e *FunctionStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitFunctionStmt(e)
}


type IfStmt struct {
    Condition Expr
	Then Stmt
	Else Stmt
}

func (e *IfStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitIfStmt(e)
}

func (e *IfStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitIfStmt(e)
}

func (e *IfStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitIfStmt(e)
}


type WhileStmt struct {
    Condition Expr
	Statement Stmt
}

func (e *WhileStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitWhileStmt(e)
}

func (e *WhileStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitWhileStmt(e)
}

func (e *WhileStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitWhileStmt(e)
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


type BlockStmt struct {
    Statements []Stmt
}

func (e *BlockStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitBlockStmt(e)
}

func (e *BlockStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitBlockStmt(e)
}

func (e *BlockStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitBlockStmt(e)
}


type ReturnStmt struct {
    Keyword Token
	Value Expr
}

func (e *ReturnStmt) AcceptString(visitor StmtVisitor[string]) string {
    return visitor.VisitReturnStmt(e)
}

func (e *ReturnStmt) AcceptInterface(visitor StmtVisitor[interface{}]) interface{} {
    return visitor.VisitReturnStmt(e)
}

func (e *ReturnStmt) Accept(visitor StmtVisitorVoid)  {
    visitor.VisitReturnStmt(e)
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



