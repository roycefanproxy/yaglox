package main

import "strconv"

var (
	keywords map[string]TokenType
)

func init() {
	keywords = map[string]TokenType{
		"and":    And,
		"class":  Class,
		"else":   Else,
		"false":  False,
		"for":    For,
		"func":   Func,
		"if":     If,
		"nil":    Nil,
		"or":     Or,
		"print":  Print,
		"return": Return,
		"super":  Super,
		"this":   This,
		"true":   True,
		"var":    Var,
		"while":  While,
	}
}

type Scanner struct {
	src                  string
	runes                []rune
	tokens               []Token
	start, current, line int
}

func NewTokenizer(src string) *Scanner {
	return &Scanner{
		src:     src,
		runes:   []rune(src),
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) Parse() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.parseOne()
	}

	s.tokens = append(s.tokens, NewToken(EOF, []rune{}, nil, s.line))

	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.runes)
}

func (s *Scanner) parseOne() {
	char := s.advance()

	switch char {
	case '(':
		s.addToken(LeftParen)
	case ')':
		s.addToken(RightParen)
	case '{':
		s.addToken(LeftBrace)
	case '}':
		s.addToken(RightBrace)
	case ',':
		s.addToken(Comma)
	case '.':
		s.addToken(Dot)
	case '-':
		s.addToken(Minus)
	case '+':
		s.addToken(Plus)
	case ';':
		s.addToken(Semicolon)
	case '*':
		s.addToken(Star)

	case '!':
		tokenType := Bang
		if s.match('=') {
			tokenType = BangEqual
		}
		s.addToken(tokenType)
	case '=':
		tokenType := Equal
		if s.match('=') {
			tokenType = EqualEqual
		}
		s.addToken(tokenType)
	case '<':
		tokenType := Less
		if s.match('=') {
			tokenType = LessEqual
		}
		s.addToken(tokenType)
	case '>':
		tokenType := Greater
		if s.match('=') {
			tokenType = GreaterEqual
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
			s.addToken(Slash)
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

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := string(s.runes[s.start:s.current])
	tokenType, hasVal := keywords[text]
	if !hasVal {
		tokenType = Identifier
	}

	s.addToken(tokenType)
}

func (s *Scanner) isAlphaNumeric(char rune) bool {
	return s.isAlpha(char) || s.isDigit(char)
}

func (s *Scanner) isAlpha(char rune) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_')
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
	}

	num, _ := strconv.ParseFloat(string(s.runes[s.start:s.current]), 64)
	s.addTokenWithLiteral(Number, num)
}

func (s *Scanner) isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func (s *Scanner) string() {
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
	s.addTokenWithLiteral(String, value)
}

func (s *Scanner) peekNext() rune {
	if (s.current + 1) >= len(s.runes) {
		return rune(0)
	}

	return s.runes[s.current+1]
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return rune(0)
	}

	return s.runes[s.current]
}

func (s *Scanner) match(char rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.runes[s.current] != char {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) advance() rune {
	char := s.runes[s.current]
	s.current++

	return char
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.runes[s.start:s.current]

	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}
