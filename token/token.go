package token

type TokenType string

type Token struct {
	Type 	TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT	= "IDENT" // add, foobar, x, y, ...
	INT		= "INT" 	// 123456

	//Operators
	ASSIGN	= "="
	PLUS	= "+"
	MINUS	= "-"
	BANG	= "!"
	ASTERIK	= "*"
	SLASH	= "/"

	LT		= "<"
	GT		= ">"
	EQ		= "=="
	NOT_EQ 	= "!="

	//Delimiters
	COMMA		= ","
	SEMICOLON	= ";"

	LPAREN	= "("
	RPAREN	= ")"
	LBRACE	= "{"
	RBRACE	= "}"

	//Keywords
	FUNCTION	= "FUNCTION"
	LET			= "LET"
	IF			= "IF"
	ELSE 		= "FALSE"
	TRUE		= "TRUE"
	FALSE       = "FALSE"
	RETURN 		= "RETURN"
)

var Keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}