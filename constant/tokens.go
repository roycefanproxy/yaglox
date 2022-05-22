package constant

//go:generate stringer -type=TokenType -output=tokens_string.go

type TokenType int

const (
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	Identifier
	String
	Number

	And
	Class
	Else
	False
	Func
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)
