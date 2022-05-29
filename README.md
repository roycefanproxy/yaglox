# yaglox
Yet another go lox

## Syntax

program -> declaration\* EOF;

declaration -> funcDecl | varDecl | statement;

funcDecl -> "func" function;

function -> IDENTIFIER "(" params? ")" block;

params -> IDENTIFIER ("," IDENTIFIER)*;

statement -> exprStmt | ifStmt | forStmt | whileStmt | block | returnStmt | printStmt;

exprStmt -> expression ";";

ifStmt -> "if" "(" expression ")" statement ("else" statement)?;

forStmt -> "for" "(" (varDecl | exprStmt | ";") expression? ";" expression? ")" statement;

while -> "while" "(" expression ")" statement;

block -> "{" declaration* "}"

returnStmt -> "return" expression? ";";

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

unary -> ((! | -) unary) | call;

call -> primary ("(" arguments? ")")*;

arguments -> expression ("," expression)*;

primary -> NUMBER | STRING | "false" | "true" | "nil" | grouping | IDENTIFIER;

grouping -> "(" expression ")";
