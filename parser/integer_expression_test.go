package parser

import (
	"testing"

	"github.com/yadavkl/monkey/ast"
	"github.com/yadavkl/monkey/lexer"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program not have enough statement got: %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program statement[0] is not expression statement got: %T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt expression is not interger literal got: %T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Fatalf("literal value is not %d got %d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("token.Literal is not %s got %s", "5", literal.TokenLiteral())
	}
}
