package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}
