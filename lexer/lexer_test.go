package lexer

import (
	"fmt"
	"testing"

	"github.com/oldmonad/json-parser/token"
)

func TestNextToken1(t *testing.T) {
	input := `{"test": "here"}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"test\""},
		{token.COLON, ":"},
		{token.STRING, "\"here\""},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `{test: here}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "test"},
		{token.COLON, ":"},
		{token.STRING, "here"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `{"test": "here", "mooch": 10}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"test\""},
		{token.COLON, ":"},
		{token.STRING, "\"here\""},
		{token.COMMA, ","},
		{token.STRING, "\"mooch\""},
		{token.COLON, ":"},
		{token.NUMBER, "10"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken4(t *testing.T) {
	input := `{"mooch": 10.9}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"mooch\""},
		{token.COLON, ":"},
		{token.NUMBER, "10.9"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken5(t *testing.T) {
	input := `{"bracket": []}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"bracket\""},
		{token.COLON, ":"},
		{token.LEFTBRACKET, "["},
		{token.RIGHTBRACKET, "]"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken6(t *testing.T) {
	input := `{"bracket": ["five", 5, 6.0]}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"bracket\""},
		{token.COLON, ":"},
		{token.LEFTBRACKET, "["},
		{token.STRING, "\"five\""},
		{token.COMMA, ","},
		{token.NUMBER, "5"},
		{token.COMMA, ","},
		{token.NUMBER, "6.0"},
		{token.RIGHTBRACKET, "]"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken7(t *testing.T) {
	input := `{"bracket": ["five", 5, 6.0], "curly_braces": {}}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"bracket\""},
		{token.COLON, ":"},
		{token.LEFTBRACKET, "["},
		{token.STRING, "\"five\""},
		{token.COMMA, ","},
		{token.NUMBER, "5"},
		{token.COMMA, ","},
		{token.NUMBER, "6.0"},
		{token.RIGHTBRACKET, "]"},
		{token.COMMA, ","},
		{token.STRING, "\"curly_braces\""},
		{token.COLON, ":"},
		{token.LEFTCURLYBRACE, "{"},
		{token.RIGHTCURLYBRACE, "}"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken8(t *testing.T) {
	input := `{"bracket": ["five", 5, 6.0], "curly_braces": {"boolean": true, "boolean_false": false}, "null_value": null}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"bracket\""},
		{token.COLON, ":"},
		{token.LEFTBRACKET, "["},
		{token.STRING, "\"five\""},
		{token.COMMA, ","},
		{token.NUMBER, "5"},
		{token.COMMA, ","},
		{token.NUMBER, "6.0"},
		{token.RIGHTBRACKET, "]"},
		{token.COMMA, ","},
		{token.STRING, "\"curly_braces\""},
		{token.COLON, ":"},
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"boolean\""},
		{token.COLON, ":"},
		{token.BOOLEAN, "true"},
		{token.COMMA, ","},
		{token.STRING, "\"boolean_false\""},
		{token.COLON, ":"},
		{token.BOOLEAN, "false"},
		{token.RIGHTCURLYBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "\"null_value\""},
		{token.COLON, ":"},
		{token.NULL, "null"},
		{token.RIGHTCURLYBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			fmt.Println(tt.expectedType)
			fmt.Println(tt.expectedValue)
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken9(t *testing.T) {
	input := `{"test": "here", "mooch": -10}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"test\""},
		{token.COLON, ":"},
		{token.STRING, "\"here\""},
		{token.COMMA, ","},
		{token.STRING, "\"mooch\""},
		{token.COLON, ":"},
		{token.NUMBER, "-10"},
		{token.RIGHTCURLYBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}

func TestNextToken10(t *testing.T) {
	input := `{"test": "here", "mooch": -10.5}`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LEFTCURLYBRACE, "{"},
		{token.STRING, "\"test\""},
		{token.COLON, ":"},
		{token.STRING, "\"here\""},
		{token.COMMA, ","},
		{token.STRING, "\"mooch\""},
		{token.COLON, ":"},
		{token.NUMBER, "-10.5"},
		{token.RIGHTCURLYBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}
