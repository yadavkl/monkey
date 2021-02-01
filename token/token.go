package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	//Identifiers + literals
	INDENT = "INDENT"
	INT    = "INT"
	//operators
	PLUS     = "+"
	ASSIGN   = "="
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	//delemeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	ELSE     = "ELSE"
	IF       = "IF"
	RETURN   = "RETURN"
)

var keyword = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
	"if":     IF,
	"return": RETURN,
}

func LookupIndent(indent string) TokenType {
	if tok, ok := keyword[indent]; ok {
		return tok
	}
	return INDENT
}
