package parser

import (
	"testing"

	"github.com/yadavkl/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
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

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for idx, tc := range tests {
		stmt := program.Statements[idx]
		if !testLetStatement(t, stmt, tc.expectedIdentifier) {
			return
		}
	}
}
