package structs

type Token struct {
	tokenType int
	lexeme    string
	literal   string
	line      int
}

func (t Token) ToTokenString() string {
	return string(t.tokenType) + " " + t.lexeme + " " + t.literal
}
