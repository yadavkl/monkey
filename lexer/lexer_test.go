package lexer

import (
	"testing"

	"github.com/yadavkl/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for idx, tc := range tests {
		tok := l.NextToken()
		if tok.Type != tc.expectedType {
			t.Errorf("test[%v]- token type wrong expected %q got %q\n", idx, tc.expectedType, tok.Type)
		}
		if tok.Literal != tc.expectedLiteral {
			t.Errorf("test[%v] token literal wrong expected %q got %q \n", idx, tc.expectedLiteral, tok.Literal)
		}

	}

}
