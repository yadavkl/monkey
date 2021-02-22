package parser

import (
	"testing"

	"github.com/yadavkl/monkey/ast"
	"github.com/yadavkl/monkey/lexer"
)

func TestIdenfierExpression(t *testing.T) {
	input := `foobar;`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("Program dont have enough statements, got: %v\n", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] not ast.ExpressionStatement Got: %T\n", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.Identifier Got: %T\n", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not %s Got %s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}
