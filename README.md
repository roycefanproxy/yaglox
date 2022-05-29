# yaglox
Yet another go lox

## Syntax

program -> declaration\* EOF;

declaration -> varDecl | statement;

statement -> exprStmt | ifStmt | forStmt | whileStmt | block | printStmt;

exprStmt -> expression ";";

ifStmt -> "if" "(" expression ")" statement ("else" statement)?;

forStmt -> "for" "(" (varDecl | exprStmt | ";") expression? ";" expression? ")" statement;

while -> "while" "(" expression ")" statement;

block -> "{" declaration* "}"

printStmt -> "print" expression ";";

varDecl -> "var" IDENTIFIER ("=" expression)? ";";

expression -> assignment;

assignment -> IDENTIFIER "=" assignment | or;

or -> and ("or" and)*;

and -> equality ("and" equality)*;

equality -> comparison ((!= | ==) comparison)\*;

comparison -> term ((< | <= | >= | >) term)\*;

term -> factor ((- | +) factor)\*;

factor -> unary ((/ | \*) unary)\*;

unary -> ((! | -) unary) | primary;

primary -> NUMBER | STRING | "false" | "true" | "nil" | grouping | IDENTIFIER;

grouping -> "(" expression ")";
