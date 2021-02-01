package lexer

import (
	"testing"

	"github.com/yadavkl/monkey/token"
)

// func TestNextToken(t *testing.T) {
// 	input := `=+(){},;`
// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.ASSIGN, "="},
// 		{token.PLUS, "+"},
// 		{token.LPAREN, "("},
// 		{token.RPAREN, ")"},
// 		{token.LBRACE, "{"},
// 		{token.RBRACE, "}"},
// 		{token.COMMA, ","},
// 		{token.SEMICOLON, ";"},
// 		{token.EOF, ""},
// 	}
// 	l := New(input)

// 	for idx, tc := range tests {
// 		tok := l.NextToken()
// 		if tok.Type != tc.expectedType {
// 			t.Errorf("test[%v]- token type wrong expected %q got %q\n", idx, tc.expectedType, tok.Type)
// 		}
// 		if tok.Literal != tc.expectedLiteral {
// 			t.Errorf("test[%v] token literal wrong expected %q got %q \n", idx, tc.expectedLiteral, tok.Literal)
// 		}

// 	}

// }

func TestNextToken(t *testing.T) {
	input := `let five = 5;
		let ten = 10;
		let add = fn(x, y) {
		x + y;
		};
		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;
		if (5 < 10) {
			return true;
			} else {
			return false;
		}
		10 == 10;
		10 != 9;
		`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.INDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.INDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.INDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.INDENT, "x"},
		{token.COMMA, ","},
		{token.INDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.INDENT, "x"},
		{token.PLUS, "+"},
		{token.INDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.INDENT, "result"},
		{token.ASSIGN, "="},
		{token.INDENT, "add"},
		{token.LPAREN, "("},
		{token.INDENT, "five"},
		{token.COMMA, ","},
		{token.INDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
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
