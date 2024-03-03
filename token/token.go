package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	STRING          = "STRING"
	NUMBER          = "NUMBER"
	NULL            = "NULL"
	BOOLEAN         = "BOOLEAN"
	QUOTE           = "\""
	COMMA           = ","
	LEFTCURLYBRACE  = "{"
	RIGHTCURLYBRACE = "}"
	LEFTBRACKET     = "["
	RIGHTBRACKET    = "]"
	COLON           = ":"
	EOF             = "EOF"
	ILLEGAL         = "ILLEGAL"
)

var keywords = map[string]TokenType{
	"true":  BOOLEAN,
	"false": BOOLEAN,
	"null":  NULL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return STRING
}
