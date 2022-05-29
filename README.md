# yaglox
Yet another go lox

## Syntax

program -> declaration\* EOF;

declaration -> varDecl | statement;

statement -> exprStmt | printStmt | block;

exprStmt -> expression ";";

printStmt -> "print" expression ";";

block -> "{" declaration* "}"

varDecl -> "var" IDENTIFIER ("=" expression)? ";";

expression -> assignment;

assignment -> IDENTIFIER "=" assignment | equality;

equality -> comparison ((!= | ==) comparison)\*;

comparison -> term ((< | <= | >= | >) term)\*;

term -> factor ((- | +) factor)\*;

factor -> unary ((/ | \*) unary)\*;

unary -> ((! | -) unary) | primary;

primary -> NUMBER | STRING | "false" | "true" | "nil" | grouping | IDENTIFIER;

grouping -> "(" expression ")";
