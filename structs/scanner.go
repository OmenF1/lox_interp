package structs

import tokens "lox/interpreter/constants"

type Scanner struct {
	Source string
	Tokens []Token

	start   int
	current int
	line    int
}

func (s Scanner) ScanTokens() {

	for {
		if s.isAtEnd() {
			break
		}
		s.start = s.current
		s.scanToken()
	}
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(tokens.LeftParen)
		break

	case ")":
		s.addToken(tokens.RightParen)
		break
	case "{":
		s.addToken(tokens.LeftBrace)
		break
	case "}":
		s.addToken(tokens.RightBrace)
		break

	case ",":
		s.addToken(tokens.Comma)
		break
	case ".":
		s.addToken(tokens.Dot)
		break

	case "-":
		s.addToken(tokens.Minus)
		break
	case "+":
		s.addToken(tokens.Plus)
		break
	case ";":
		s.addToken(tokens.Semicolon)
		break
	case "*":
		s.addToken(tokens.Star)
		break
	default:
		//TODO: Throw Error here.
		break
	}
}

func (s Scanner) advance() string {
	return string([]rune(s.Source)[s.current+1])
}

func (s Scanner) addToken(tokenType int) {
	s.addTokenWithLiteral(tokenType, "")
}

func (s Scanner) addTokenWithLiteral(tokenType int, literal string) {
	text := string(s.Source[s.start:s.current])
	_token := Token{
		tokenType: tokenType,
		lexeme:    text,
		literal:   literal,
		line:      s.line,
	}
	s.Tokens = append(s.Tokens, _token)
}
