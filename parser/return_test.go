package parser

import (
	"testing"

	"github.com/yadavkl/monkey/ast"
	"github.com/yadavkl/monkey/lexer"
)

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParsePrgram() return nil\n")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program does not have 3 statements got %v\n", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not returnstatement, got %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return' got %q\n", returnStmt.TokenLiteral())
		}
	}
}
