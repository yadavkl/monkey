package ast

import (
	"testing"

	"github.com/yadavkl/monkey/token"
)

func TetsString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.INDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.INDENT, Literal: "anotherVal"},
					Value: "anotherVal",
				},
			},
		},
	}
	if program.String() != "let myVal = anotherVal;" {
		t.Errorf("Expected: 'let myVar = anotherVal' Got: %s", program.String())
	}

}
