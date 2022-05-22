# yaglox
Yet another go lox

## Syntax

program -> statement\* EOF;

statement -> exprStmt | printStmt;

exprStmt -> expression ";";

printStmt -> "print" expression ";";

expression -> equality;

equality -> comparison ((!= | ==) comparison)\*;

comparison -> term ((< | <= | >= | >) term)\*;

term -> factor ((- | +) factor)\*;

factor -> unary ((/ | \*) unary)\*;

unary -> ((! | -) unary) | primary;

primary -> NUMBER | STRING | "false" | "true" | "nil" | grouping;

grouping -> "(" expression ")";
