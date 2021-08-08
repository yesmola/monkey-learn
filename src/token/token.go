package token

type Type string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// Identifiers and literals
const (
	IDENT = "IDENT" // identifiers add, foobar, x, y, ...
	INT   = "INT" // int literals
)

// Operators
const (
	ASSIGN = "="
	PLUS   = "+"
)

// Delimiters
const (
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

// Keywords
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    Type
	Literal string
}