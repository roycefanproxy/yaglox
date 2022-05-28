package main

import (
	"strconv"

	"github.com/roycefanproxy/yaglox/constant"
)

var (
	keywords map[string]constant.TokenType
)

func init() {
	keywords = map[string]constant.TokenType{
		"and":    constant.And,
		"class":  constant.Class,
		"else":   constant.Else,
		"false":  constant.False,
		"for":    constant.For,
		"func":   constant.Func,
		"if":     constant.If,
		"nil":    constant.Nil,
		"or":     constant.Or,
		"print":  constant.Print,
		"return": constant.Return,
		"super":  constant.Super,
		"this":   constant.This,
		"true":   constant.True,
		"var":    constant.Var,
		"while":  constant.While,
	}
}

type Tokenizer struct {
	src                  string
	runes                []rune
	tokens               []Token
	start, current, line int
}

func NewTokenizer(src string) *Tokenizer {
	return &Tokenizer{
		src:     src,
		runes:   []rune(src),
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Tokenizer) Parse() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.parseOne()
	}

	s.tokens = append(s.tokens, NewToken(constant.EOF, []rune{}, nil, s.line))

	return s.tokens
}

func (s *Tokenizer) isAtEnd() bool {
	return s.current >= len(s.runes)
}

func (s *Tokenizer) parseOne() {
	char := s.advance()

	switch char {
	case '(':
		s.addToken(constant.LeftParen)
	case ')':
		s.addToken(constant.RightParen)
	case '{':
		s.addToken(constant.LeftBrace)
	case '}':
		s.addToken(constant.RightBrace)
	case ',':
		s.addToken(constant.Comma)
	case '.':
		s.addToken(constant.Dot)
	case '-':
		s.addToken(constant.Minus)
	case '+':
		s.addToken(constant.Plus)
	case ';':
		s.addToken(constant.Semicolon)
	case '*':
		s.addToken(constant.Star)

	case '!':
		tokenType := constant.Bang
		if s.match('=') {
			tokenType = constant.BangEqual
		}
		s.addToken(tokenType)
	case '=':
		tokenType := constant.Equal
		if s.match('=') {
			tokenType = constant.EqualEqual
		}
		s.addToken(tokenType)
	case '<':
		tokenType := constant.Less
		if s.match('=') {
			tokenType = constant.LessEqual
		}
		s.addToken(tokenType)
	case '>':
		tokenType := constant.Greater
		if s.match('=') {
			tokenType = constant.GreaterEqual
		}
		s.addToken(tokenType)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else if s.match('*') {
			for s.peek() != '*' && s.peekNext() != '/' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(constant.Slash)
		}
	case ' ', '\r', '\t':
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(char) {
			s.number()
		} else if s.isAlpha(char) {
			s.identifier()
		} else {
			EmitErrorLog(s.line, "Unexpected character.")
		}
	}
}

func (s *Tokenizer) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := string(s.runes[s.start:s.current])
	tokenType, hasVal := keywords[text]
	if !hasVal {
		tokenType = constant.Identifier
	}

	s.addToken(tokenType)
}

func (s *Tokenizer) isAlphaNumeric(char rune) bool {
	return s.isAlpha(char) || s.isDigit(char)
}

func (s *Tokenizer) isAlpha(char rune) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_')
}

func (s *Tokenizer) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
	}

	for s.isDigit(s.peek()) {
		s.advance()
	}

	num, _ := strconv.ParseFloat(string(s.runes[s.start:s.current]), 64)
	s.addTokenWithLiteral(constant.Number, num)
}

func (s *Tokenizer) isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func (s *Tokenizer) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		EmitErrorLog(s.line, "Unterminated string.")
		return
	}

	s.advance()

	value := string(s.runes[s.start+1 : s.current-1])
	s.addTokenWithLiteral(constant.String, value)
}

func (s *Tokenizer) peekNext() rune {
	if (s.current + 1) >= len(s.runes) {
		return rune(0)
	}

	return s.runes[s.current+1]
}

func (s *Tokenizer) peek() rune {
	if s.isAtEnd() {
		return rune(0)
	}

	return s.runes[s.current]
}

func (s *Tokenizer) match(char rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.runes[s.current] != char {
		return false
	}

	s.current++

	return true
}

func (s *Tokenizer) advance() rune {
	char := s.runes[s.current]
	s.current++

	return char
}

func (s *Tokenizer) addToken(tokenType constant.TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Tokenizer) addTokenWithLiteral(tokenType constant.TokenType, literal interface{}) {
	text := s.runes[s.start:s.current]

	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}
