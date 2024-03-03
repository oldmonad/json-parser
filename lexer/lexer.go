package lexer

import (
	"github.com/oldmonad/json-parser/token"
)

type Lexer struct {
	source          string
	currentPosition int
	readPosition    int
	character       byte
	line            int
}

func New(source string) *Lexer {
	l := &Lexer{source: source}
	l.line += 1
	l.setCurrentCharacter()
	return l
}

func (l *Lexer) setCurrentCharacter() {
	if isAtEnd(l) {
		l.character = 0
	} else {
		l.character = l.source[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '{':
		tok = newToken(token.LEFTCURLYBRACE, l.character)
	case '}':
		tok = newToken(token.RIGHTCURLYBRACE, l.character)
	case '[':
		tok = newToken(token.LEFTBRACKET, l.character)
	case ']':
		tok = newToken(token.RIGHTBRACKET, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case ':':
		tok = newToken(token.COLON, l.character)
	case '\n':
		l.line += 1
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	case '-':
		if isDigit(l.peekNextCharacter()) {
			ch := l.character
			l.setCurrentCharacter()
			tok = token.Token{Type: token.NUMBER, Value: string(ch) + string(l.readNumber())}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.character)
		}
	default:
		if isLetter(l.character) {
			tok.Value = l.readString()
			tok.Type = token.LookupIdent(tok.Value)
			return tok
		} else if isDigit(l.character) {
			tok.Type = token.NUMBER
			tok.Value = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.character)
		}
	}

	l.setCurrentCharacter()

	return tok
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition

	for isDigit(l.character) {
		l.setCurrentCharacter()
	}

	if l.character == '.' {
		l.setCurrentCharacter()

		for isDigit(l.character) {
			l.setCurrentCharacter()
		}
	}

	return l.source[position:l.currentPosition]
}

func (l *Lexer) readString() string {
	currentPosition := l.currentPosition

	for isLetter(l.character) {
		l.setCurrentCharacter()
	}

	return l.source[currentPosition:l.currentPosition]
}

func (l *Lexer) peekNextCharacter() byte {
	if l.readPosition >= len(l.source) {
		return 0
	} else {
		return l.source[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Value: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '"'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAtEnd(l *Lexer) bool {
	return l.readPosition >= len(l.source)
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\r' {
		l.setCurrentCharacter()
	}
}
