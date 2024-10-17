package structs

import tokens "lox/interpreter/constants"

type Scanner struct {
	Source string
	Tokens []Token

	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() {

	for {
		if s.isAtEnd() {
			break
		}
		s.start = s.current
		s.scanToken()
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(tokens.LeftParen)

	case ")":
		s.addToken(tokens.RightParen)

	case "{":
		s.addToken(tokens.LeftBrace)

	case "}":
		s.addToken(tokens.RightBrace)

	case ",":
		s.addToken(tokens.Comma)

	case ".":
		s.addToken(tokens.Dot)

	case "-":
		s.addToken(tokens.Minus)

	case "+":
		s.addToken(tokens.Plus)

	case ";":
		s.addToken(tokens.Semicolon)

	case "*":
		s.addToken(tokens.Star)

	case "!":
		//	So go doesn't have ternary operations?
		if s.match("=") {
			s.addToken(tokens.BangEqual)
		} else {
			s.addToken(tokens.Bang)
		}

	default:
		//TODO: Throw Error here.

	}
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}

	if string(s.Source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) advance() string {
	return string(s.Source[s.current+1])
	// return string([]rune(s.Source)[s.current++])
}

func (s *Scanner) addToken(tokenType int) {
	s.addTokenWithLiteral(tokenType, "")
}

func (s *Scanner) addTokenWithLiteral(tokenType int, literal string) {
	text := string(s.Source[s.start:s.current])
	_token := Token{
		tokenType: tokenType,
		lexeme:    text,
		literal:   literal,
		line:      s.line,
	}
	s.Tokens = append(s.Tokens, _token)
}
